import { is_available, token, messages } from "@stores/chat";
import { GetChatCredentialsRequest, GetChatRoomRequest, PostChatMessageRequest, GetChatCredentialsVerificationRequest } from "@libs/HttpRequests";
import { get } from "svelte/store";

export class ChatMessage {
    constructor({ chat_id, message, sender, send_date }) {
        this.chat_id = chat_id;
        this.content = message;
        this.sender = sender;
        this.send_date = send_date;
    }
}

class ChatNODE {
    constructor (message, back, next) {
        this.message = message;
        this.back = back;
        this.next = next;
    }
}

export class ChatQueue {
    
    /**
     * @type {ChatNODE}
    */

    #head;
    /**
     * @type {ChatNODE}
    */
    
    #tail;
    /**
     * @type {ChatNODE}
     */
    #traverser;
    constructor() {
        this.#head = null;
        this.#tail = null;
        this.#traverser = null;
    }

    enqueue = (message) => {
        let node = new ChatNODE(message, null, this.#tail);

        if (this.#tail) {
            this.#tail.back = node;
        }

        this.#tail = node;
        this.#head = this.#head || node;
        
    }

    dequeue = () => {
        let node = this.#head;
        this.#head = this.#head?.back;

        if (this.#head) {
            this.#head.next = null;
        } else {
            this.#tail = null;
        }

        return node?.message;
    }

    
    /**
     * Yields the next message in the queue
     * @date 9/11/2023 - 8:26:16 PM
     * @returns {ChatMessage}
    */
    traverse = () => {
        this.#traverser = this.#traverser?.back || this.#head;
        return this.#traverser?.message;
    }

    isEmpty = () => {
        return this.#head === null;
    }

}

export class ChatDialer {
    constructor() {
        this.messages_queue = new ChatQueue();
        this.chat_id = null;
        this.chat_ready = get(token) !== "";

        this.init();
    }

    init = async () => {
        if (!this.chat_ready) {
            let are_claims_valid = await this.#verifyClaims();
            console.log(`Claims are valid: ${are_claims_valid}`);
            if (are_claims_valid) {
                this.#requestChatRoom();
            }
        }
    }

    
    /**
     * Sends a message to the portfolio service
     * @date 9/12/2023 - 9:21:29 PM
     *
     * @async
     * @param {ChatMessage} message
     * @returns {ChatMessage}
     */
    #dialMessage = async (message) => {
        const message_promise = new Promise((resolve, reject) => {
            const message_request = new PostChatMessageRequest({ chat_id: this.chat_id, message: message.content});

            const on_success = message_data => {
                let new_message = new ChatMessage({
                    chat_id: this.chat_id,
                    message: message_data.content,
                    sender: message_data.author,
                    send_date: message_data.send_date
                });

                resolve(new_message);
            }

            const on_error = (status) => {
                console.log(`Error ${status} while posting chat message`);
                reject(new Error(`Error ${status} while posting chat message`));
            }

            message_request.do(on_success, on_error);
        });

        let new_message = await message_promise.catch((error) => {
            console.log(error);
            return null;
        });

        if (new_message !== null) {
            this.messages_queue.enqueue(new_message);

            const visible_messages = get(messages);
            visible_messages.push(new_message);
            
            messages.set(visible_messages);
        }

        return new_message;
    }

    #requestChatRoom = async () => {
        const chat_room_promise = new Promise((resolve, reject) => {
            const chat_room_request = new GetChatRoomRequest();

            const on_success = (chat_room) => {
                resolve(chat_room);
            }

            const on_error = (status) => {
                console.log(`Error ${status} while getting chat room`);
                reject(new Error(`Error ${status} while getting chat room`));
            }

            chat_room_request.do(on_success, on_error);
        });

        const chat_room = await chat_room_promise.catch((error) => console.log(error));

        if (chat_room) {
            this.chat_id = chat_room.id;
            const visible_messages = [];
            chat_room.messages.forEach((message, h) => {
                let new_message = new ChatMessage({ 
                    chat_id: this.chat_id,
                    message: message.content,
                    sender: message.author,
                    send_date: message.send_date
                })

                this.messages_queue.enqueue(new_message);

                if (new_message.sender !== "system") {
                    visible_messages.push(new_message);
                }
            });
            console.log(visible_messages);
            messages.set(visible_messages);
        }


    }

    #setClaims = async () => {
        let claims_promised = new Promise((resolve, reject) => {
            const claims_request = new GetChatCredentialsRequest();
            console.log("Getting chat credentials");

            const on_success = (token) => {
                console.log("Chat credentials received");
                resolve(token);
            }

            const on_error = (status) => {
                reject(new Error(`Error ${status} while getting chat credentials`));
            }

            claims_request.do(on_success, on_error);
        });

        let claims = await claims_promised.catch((error) => {
            is_available.set(false);
            token.set("");
            console.log(`Error: ${error}`);
            return null;
        });
        
        if (claims !== null) {
            console.log("Claims set: "+JSON.stringify(claims));
            this.chat_ready = true;
            token.set(claims);
            return true;
        }

        return false;
    }

    sendMessage = (message_content, send=true) => {
        let message = new ChatMessage({
            chat_id: this.chat_id,
            message: message_content,
            sender: "user",
            send_date: new Date().toUTCString()
        });

        if (this.chat_ready) {
            this.messages_queue.enqueue(message);
            const visible_messages = get(messages);
            visible_messages.push(message);

            
            if (send) {
                this.#dialMessage(message);
            }

            messages.set(visible_messages);
            
        } else {
            console.log("Chat is not ready");
        }
    }

    #verifyClaims = async () => {
        let claims_promised = new Promise((resolve, reject) => {
            const claims_request = new GetChatCredentialsVerificationRequest();

            const on_success = (token) => {
                resolve(token);
            }

            const on_error = (status) => {
                reject(new Error(`Error ${status} while validating credentials, reauthenticating`));
            }

            claims_request.do(on_success, on_error);
        });

        let claims = await claims_promised.catch((error) => {
            console.log(error);
            return null;
        });

        console.log(`claims: ${JSON.stringify(claims)}`);

        if (claims !== null) {
            console.log("Claims verified");
            this.chat_ready = true;
            token.set(claims);
            return true;
        }

        return this.#setClaims();
    }
}

export function testQueue() {
    let queue = new ChatQueue();

    queue.enqueue(new ChatMessage({ chat_id: 1, message: "Hello A", sender: "user", order: 1 }));
    queue.enqueue(new ChatMessage({ chat_id: 1, message: "Hello B", sender: "user", order: 2 }));
    queue.enqueue(new ChatMessage({ chat_id: 1, message: "Hello C", sender: "user", order: 3 }));

    console.log(queue.traverse());
    console.log(queue.traverse());
    console.log(queue.traverse());

    console.log(queue.dequeue());
    console.log(queue.dequeue());
    console.log(queue.dequeue());

    console.log(queue.isEmpty());

    return queue;
}
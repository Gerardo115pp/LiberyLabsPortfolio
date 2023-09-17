export const portfolio_service = PORTFOLIO_SERVICE;
window.portfolio_service = portfolio_service;

function attributesToJson() {
    const json_data = {};

    Object.entries(this).forEach(([key, value]) => {
        if(!(this[key] instanceof Function) && key[0] !== "_") {
            json_data[key] = value;
        }
    });

    return JSON.stringify(json_data);
}

export class GetAllProjectsRequest {
    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');

        const request = new Request(`${portfolio_service}/projects`, {
            method: 'GET',
            headers: headers
        });

        fetch(request).then(promise => {
            if (promise.status >= 200 && promise.status < 300) {
                promise.json().then(data => {
                    on_success(data);
                })
            } else {
                on_error(promise.status);
            }
        });
    }
}

export class GetProjectById {
    constructor(id) {
        this.id = id;
    }

    toJson =attributesToJson.bind(this);

    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');

        const request = new Request(`${portfolio_service}/projects?id=${this.id}`, {
            method: 'GET',
            headers: headers
        });

        fetch(request).then(promise => {
            if (promise.status >= 200 && promise.status < 300) {
                promise.json().then(data => {
                    on_success(data);
                })
            } else {
                on_error(promise.status);
            }
        });
    }
}

export class GetProjectImagesData {
    constructor(id) {
        this.id = id;
    }

    toJson = attributesToJson.bind(this);


    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');

        const request = new Request(`${portfolio_service}/project-images?project_id=${this.id}`, {
            method: 'GET',
            headers: headers
        });

        fetch(request).then(promise => {
            if (promise.status >= 200 && promise.status < 300) {
                promise.json().then(data => {
                    on_success(data);
                })
            } else {
                on_error(promise.status);
            }
        });
    }
}

export class GetProjectIdeasRequest {
    toJson = attributesToJson.bind(this);

    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');

        const request = new Request(`${portfolio_service}/project-ideas`, {
            method: 'GET',
            headers: headers
        });

        fetch(request).then(promise => {
            if (promise.status >= 200 && promise.status < 300) {
                promise.json().then(data => {
                    on_success(data);
                })
            } else {
                on_error(promise.status);
            }
        });
    }

}

export class PostContactMessageRequest {
    constructor({email, name, budget, company, start_date, end_date, subject}) {
        this.email = email;
        this.name = name;
        this.budget = budget;
        this.company = company;
        this.start_date = start_date;
        this.end_date = end_date;
        this.subject = subject;
    }

    toJson = attributesToJson.bind(this);

    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');

        const request = new Request(`${portfolio_service}/contact`, {
            method: 'POST',
            headers: headers,
            body: this.toJson()
        });

        fetch(request).then(promise => {
            if (promise.status >= 200 && promise.status < 300) {
                on_success();
            } else {
                on_error(promise.status);
            }
        });
    }
}

export class GetChatCredentialsRequest {
    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');

        const request = new Request(`${portfolio_service}/chat-claims`, {
            method: 'GET',
            headers: headers, 
            credentials: 'same-origin'
        });

        fetch(request).then(promise => {
            if (promise.status >= 200 && promise.status < 300) {
                let token = promise.headers.get('X-Chat-Token');
                return on_success(token);
            } else {
                on_error(promise.status);
            }
        });
    }
}

export class GetChatCredentialsVerificationRequest {
    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');

        const request = new Request(`${portfolio_service}/chat-claims/verify`, {
            method: 'GET',
            headers: headers, 
            credentials: 'same-origin'
        });

        fetch(request).then(promise => {
            if (promise.status >= 200 && promise.status < 300) {
                let token = promise.headers.get('X-Chat-Token');
                return on_success(token);
            } else {
                on_error(promise.status);
            }
        });
    }
}

export class GetChatRoomRequest {
    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');

        const request = new Request(`${portfolio_service}/chat`, {
            method: 'GET',
            headers: headers, 
            credentials: 'same-origin'
        });

        fetch(request).then(promise => {
            if (promise.status >= 200 && promise.status < 300) {
                promise.json().then(data => {
                    on_success(data);
                });
            } else {
                on_error(promise.status);
            }
        });
    }

}

export class PostChatMessageRequest {
    constructor({chat_id, message}) {
        this.chat_id = chat_id;
        this.content = message;
    }

    toJson = attributesToJson.bind(this);

    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');

        const request = new Request(`${portfolio_service}/chat`, {
            method: 'POST',
            headers: headers,
            body: this.toJson(),
            credentials: 'same-origin'
        });

        fetch(request).then(promise => {
            if (promise.status >= 200 && promise.status < 300) {
                promise.json().then(new_messages => {
                    on_success(new_messages);
                });
            } else {
                on_error(promise.status);
            }
        });
    }
}

export class GetRecaptchaVerificationRequest {
    constructor(token) {
        this.token = token;
    }

    toJson = attributesToJson.bind(this);

    do = (on_success, on_error) => {
        const headers = new Headers();
        headers.append('Content-Type', 'application/json');

        const request = new Request(`${portfolio_service}/is-user-human?token=${this.token}`, {
            method: 'GET',
            headers: headers,
            credentials: 'same-origin'
        });

        fetch(request).then(promise => {
            if (promise.status >= 200 && promise.status < 300) {
                promise.json().then(data => {
                    const { IsHuman:is_human } = data;
                    if(is_human) {
                        on_success();
                    } else {
                        on_error(-1);
                    }
                });
            } else {
                on_error(promise.status);
            }
        });
    }
}
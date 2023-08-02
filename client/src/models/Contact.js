import { PostContactMessageRequest } from "@libs/HttpRequests";

class ContactMessage extends PostContactMessageRequest {
    
    send = () => {
        return new Promise((resolve, reject) => {
            this.do((data) => {
                resolve(true);
            }, (error) => {
                reject(false);
            });
        });
    }
}

export default ContactMessage;
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

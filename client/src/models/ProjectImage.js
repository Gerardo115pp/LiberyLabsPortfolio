import { GetProjectImagesData, portfolio_service } from "@libs/HttpRequests";

export class ProjectImages {
    #project_id;
    #mobile_images;
    #current_mobile_image;
    #desktop_images;
    #current_desktop_image;
    #total_images;
    constructor(project_id) {
        this.#project_id = project_id;
        this.#mobile_images = [];
        this.#current_mobile_image = -1;

        this.#desktop_images = [];
        this.#current_desktop_image = -1;
    }

    loadProjectImages = () => {
        const request = new GetProjectImagesData(this.#project_id);
        const get_images_promise = new Promise((resolve, reject) => {
            const on_success = (data) => {
                this.#total_images = data["desktop_count"] + data["mobile_count"];
                this.#mobile_images = data["images"]["mobile"];
                this.#current_mobile_image = this.#mobile_images.length > 0 ? 0 : -1;
        
                this.#desktop_images = data["images"]["desktop"];
                this.#current_desktop_image = this.#desktop_images.length > 0 ? 0 : -1;
            
                resolve(true);
            }
            const on_error = (status) => reject(status);

            request.do(on_success, on_error);
        })

        return get_images_promise;
    }

    #getImageHref = image_name => `${portfolio_service}/project-images?project_id=${this.#project_id}&name=${image_name}`;

    get TotalImages() {
        return this.#total_images;
    }

    get NextMobileImage() {
        if (this.#current_mobile_image === -1) return null;

        const next_image = this.#mobile_images[this.#current_mobile_image];
        this.#current_mobile_image = (this.#current_mobile_image + 1) % this.#mobile_images.length;

        return this.#getImageHref(next_image);
    }

    get NextDesktopImage() {
        if (this.#current_desktop_image === -1) return null;

        const next_image = this.#desktop_images[this.#current_desktop_image];
        this.#current_desktop_image = (this.#current_desktop_image + 1) % this.#desktop_images.length;

        return this.#getImageHref(next_image);
    }
}
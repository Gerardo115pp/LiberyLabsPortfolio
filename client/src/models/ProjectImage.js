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
                this.#desktop_images = data["images"]["desktop"];
                this.resetImageCounter();
            
                resolve(true);
            }
            const on_error = (status) => reject(status);

            request.do(on_success, on_error);
        })

        return get_images_promise;
    }

    resetImageCounter = () => {
        this.#current_mobile_image = this.#mobile_images.length > 0 ? 0 : -1;
        this.#current_desktop_image = this.#desktop_images.length > 0 ? 0 : -1;
    }

    #getImageHref = image_name => `${portfolio_service}/project-images?project_id=${this.#project_id}&name=${image_name}`;

    get TotalImages() {
        return this.#total_images;
    }

    get NextMobileImage() {
        if (this.#current_mobile_image === -1) return null;

        this.#current_mobile_image = (this.#current_mobile_image + 1) % this.#mobile_images.length;
        const next_image = this.#mobile_images[this.#current_mobile_image];

        return this.#getImageHref(next_image);
    }

    get PreviousMobileImage() {
        if (this.#current_mobile_image === -1) return null;


        this.#current_mobile_image = (this.#current_mobile_image - 1) >= 0 ? this.#current_mobile_image - 1 : this.#mobile_images.length - 1;
        const previous_image = this.#mobile_images[this.#current_mobile_image];
        

        return this.#getImageHref(previous_image);
    }

    get CurrentMobileImage() {
        if (this.#current_mobile_image === -1) return null;

        return this.#getImageHref(this.#mobile_images[this.#current_mobile_image]);
    }


    get NextDesktopImage() {
        if (this.#current_desktop_image === -1) return null;

        this.#current_desktop_image = (this.#current_desktop_image + 1) % this.#desktop_images.length;
        const next_image = this.#desktop_images[this.#current_desktop_image];

        return this.#getImageHref(next_image);
    }

    get PreviousDesktopImage() {
        if (this.#current_desktop_image === -1) return null;

        this.#current_desktop_image = (this.#current_desktop_image - 1) >= 0 ? this.#current_desktop_image - 1 : this.#desktop_images.length - 1;
        const previous_image = this.#desktop_images[this.#current_desktop_image];

        return this.#getImageHref(previous_image);
    }

    get CurrentDesktopImage() {
        if (this.#current_desktop_image === -1) return null;

        return this.#getImageHref(this.#desktop_images[this.#current_desktop_image]);
    }

}
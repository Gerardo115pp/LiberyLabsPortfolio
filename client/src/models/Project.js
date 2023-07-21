import { GetAllProjectsRequest, GetProjectById } from "@libs/HttpRequests";

export class ProjectListItem {
    #id;
    constructor({id, name, relevance, tech_stack, type}) {
        this.#id = id;
        this.name = name;
        this.relevance = relevance;
        this.tech_stack = tech_stack;
        this.type = type;
    }

    get ID() {
        return this.#id;
    }
}

export class Project {
    #id;
    constructor({id, name, tech_stack, relevance, type, description, url, github}) {
        this.#id = id;
        this.name = name;
        this.tech_stack = tech_stack;
        this.relevance = relevance;
        this.type = type;
        this.description = description;
        this.url = url;
        this.github = github;
    }

    get ID() {
        return this.#id;
    }

    get Paragraphs() {
        const paragraphs = this.description.split("\n");
        return paragraphs;
    }
}

/**
 * Requests all projects from the portfolio service and sends them to the success_callback.
 * @date 7/20/2023 - 7:24:01 PM
 *
 * @param {function(ProjectListItem[])} success_callback
 * @param {function(number)} error_callback
 */
export function getAllProjects(success_callback, error_callback) {
    const request = new GetAllProjectsRequest();

    const on_success = (data) => {
        const projects = data.map(project => new ProjectListItem(project));
        success_callback(projects);
    }

    request.do(on_success, error_callback);
}

/**
 * Requests a project by its ID from the portfolio service and sends it to the success_callback.
 * @date 7/20/2023 - 7:24:01 PM
 * 
 * @param {string} project_id
 * @param {function(Project)} success_callback
 * @param {function(number)} error_callback
 */
export function getProjectByID(project_id, success_callback, error_callback) {
    const request = new GetProjectById(project_id);

    const on_success = (data) => {
        const project = new Project(data);
        success_callback(project);
    }

    request.do(on_success, error_callback);
}
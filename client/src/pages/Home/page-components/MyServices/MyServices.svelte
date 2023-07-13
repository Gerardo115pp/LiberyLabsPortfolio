<script>
    import ServicesHeader from "./sub-components/services_header.svelte";
    import ServiceDescriptor from "./sub-components/ServiceDescriptor/ServiceDescriptor.svelte";
    import ServicesContainer from "./sub-components/ServicesContainer/ServicesContainer.svelte";
    import { onMount, onDestroy } from "svelte";

    onMount(() => {
        service_rotation_interval = setInterval(rotateServices, 5500);
    })

    onDestroy(() => {
        clearInterval(service_rotation_interval);
    });

    let service_rotation_interval;

    // @typedef {Object} Service
    // @property {string} title
    // @property {string} description

    // @type {Service[]}
    const services = [
        {
            title: "Bespoke Web App Architecture & Design",
            description: "We conceptualize and design the blueprint of your software solution, establishing the foundational structure that ensures growth and scalability. This involves creating wireframes, defining user flows, and setting up the perfect software architecture."
        },
        {
            title: "Full-stack Web App Development",
            description: "Once the blueprint is ready, we bring your application to life. We take care of frontend development to ensure a user-friendly interface, backend services for robust functionality, and reliable database systems for secure data storage. We utilize technologies like Go, Python, Svelte, and JavaScript to deliver a seamless, high-performing product."
        },
        {
            title: "Web App Deployment and Maintenance",
            description: "We deploy your web app to the cloud and make sure that it is always up and running. We also provide maintenance services to ensure that your web app is always up-to-date and secure."
        }, 
        {
            title: "AI and Machine Learning Solutions",
            description: "Stay ahead with our AI integration services. We can build in conversational AI capabilities using tools like ChatGPT, providing your users with a dynamic, engaging experience."
        },
        {
            title: "Brand Website Design and Development",
            description: "We design and develop your brand website to help you establish your online presence. We make sure that your brand website is fast, secure, and mobile-friendly."
        },
        {
            title: "Workflow Automation",
            description: "We help you automate your business processes to help you save time and money. We also provide consulting services to help you understand how workflow automation can help your business."
        },
        {
            title: "Web Scraping and Data Extraction",
            description: "Harness the power of data with our scraping services. We automate API data extractions and create bots to gather the data you need. Whether for market research or data-driven strategies, we've got you covered."
        },
        {
            title: "Data Pipelines",
            description: "We build powerful data pipelines that handle complex extraction, transformation, and storage workflows. With our data solutions, you can unlock new insights and drive business growth."
        }
    ]
    let service_selected_index = 0;

    // @type {Service}
    let service_selected;
    // @type {number}
    
    
    /*=============================================
    =            Pointing animation            =
    =============================================*/
    
        let point_direction = 0;
        let user_interaction = false;

        const pointing_directions = {
            LEFT: -1,
            CENTER: 0,
            RIGHT: 1
        }

    
    /*=====  End of Pointing animation  ======*/
    
    

    $: if (service_selected_index >= 0 && service_selected_index < services.length) {
        service_selected = services[service_selected_index];
        point_direction = getPointingDirection(service_selected_index);
    }

    const getPointingDirection = index => {
        const column_count = 3;

        const element_position = (index+1) % column_count;

        if (element_position === 0) {
            return pointing_directions.LEFT;
        } else if (element_position === 1) {
            return pointing_directions.RIGHT;
        }

        return pointing_directions.CENTER;
    }

    const rotateServices = () => {
        if (user_interaction) return;

        service_selected_index = (service_selected_index + 1) % services.length;
        console.log(`rotating services: ${service_selected_index}`);
    }

    
</script>

<article id="my-services-section" class="section-content-layout">
    <ServicesHeader />
    <ServiceDescriptor {point_direction} service_description={service_selected?.description}/>
    <ServicesContainer bind:user_clicked={user_interaction} {services} bind:service_selected={service_selected_index}/>
</article>

<style>
    #my-services-section {
        row-gap: var(--vspacing-3);
    }
</style>
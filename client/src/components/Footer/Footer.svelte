<script>
  import MainLogo from "@components/UI/MainLogo.svelte";
  import LiberyBorderStar from "@components/UI/LiberyBorderStar.svelte";

  const home_links = [
    {
      name: "Beginning",
      link: "#hero-section"
    },
    {
      name: "About Me",
      link: "#about-me-section"
    },
    {
      name: "My Arsenal(Skills)",
      link: "#tech-stack-section"
    },
    {
      name: "Services",
      link: "#my-services-section"
    },
    {
      name: "Portfolio",
      link: "#recent-work-section"
    },
    {
      name: "Contact",
      link: "#contact-section"
    },
    {
      name: "A.I Chat",
      link: "#chat-section"
    }
  ]
</script>

<footer id="libery-labs-footer" >
    <div id="main-icon-wrapper">
      <MainLogo
        subheadline_color="var(--grey-3)"
      />
    </div>
    <div id="footer-libery-labs-star">
      <LiberyBorderStar
        stroke_color="var(--grey-2)"
      />
    </div>
    <div id="home-links-wrapper">
      <h3 id="link-parent-name">Home</h3>
      <ul id="home-links">
        {#each home_links as link_element}
          <li class="home-link">
            <a href="{link_element.link}">
              {link_element.name}
            </a>
          </li>
        {/each}
      </ul>
    </div>
    <div id="socials-wrapper">
      <a href="https://github.com/Gerardo115pp" id="github-profile-link">
        <svg viewBox="0 0 35 35" fill="none">
          <path d="M17.3404 0.708008C7.90046 0.708008 0.246216 8.56067 0.246216 18.2372C0.246216 25.983 5.14554 32.546 11.9374 34.8659C12.0327 34.8866 12.1299 34.8969 12.2274 34.8965C12.8608 34.8965 13.105 34.4309 13.105 34.0265C13.105 33.6068 13.0898 32.5078 13.0821 31.0426C12.5167 31.1752 11.9382 31.2443 11.3575 31.2487C8.06835 31.2487 7.32048 28.6922 7.32048 28.6922C6.54208 26.6699 5.42027 26.128 5.42027 26.128C3.93216 25.0825 5.41264 25.052 5.52711 25.052H5.53474C7.25179 25.2046 8.15229 26.8683 8.15229 26.8683C9.00701 28.364 10.1517 28.7837 11.1743 28.7837C11.8505 28.7702 12.5162 28.6142 13.1279 28.3259C13.2806 27.1964 13.7232 26.4257 14.2116 25.983C10.4188 25.5404 6.42761 24.037 6.42761 17.3215C6.42761 15.406 7.09154 13.8416 8.18282 12.6205C8.0073 12.1779 7.41968 10.3922 8.35071 7.98068C8.47559 7.9508 8.60395 7.93796 8.73228 7.94253C9.35042 7.94253 10.747 8.1791 13.0516 9.78168C15.852 8.99818 18.8136 8.99818 21.614 9.78168C23.9187 8.1791 25.3152 7.94253 25.9333 7.94253C26.0617 7.93796 26.19 7.9508 26.3149 7.98068C27.2459 10.3922 26.6583 12.1779 26.4828 12.6205C27.5741 13.8492 28.238 15.4136 28.238 17.3215C28.238 24.0523 24.2392 25.5328 20.4311 25.9678C21.0416 26.5096 21.5911 27.578 21.5911 29.2111C21.5911 31.5539 21.5682 33.4465 21.5682 34.0188C21.5682 34.4309 21.8048 34.8965 22.4382 34.8965C22.5407 34.8969 22.643 34.8867 22.7434 34.8659C29.543 32.546 34.4347 25.9754 34.4347 18.2372C34.4347 8.56067 26.7804 0.708008 17.3404 0.708008Z" fill="#E8E8E8"/>
        </svg>
      </a>
      <a href="https://www.linkedin.com/in/gerardo-rodriguez-sanchez-3b0024192/" id="linkedin-profile-link">
        <svg viewBox="0 0 35 35" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M31.983 0.708008H3.45021C1.89112 0.708008 0.528931 1.82981 0.528931 3.37058V31.9667C0.528931 33.5159 1.89112 34.8964 3.45021 34.8964H31.9746C33.5421 34.8964 34.7173 33.5067 34.7173 31.9667V3.37058C34.7265 1.82981 33.5421 0.708008 31.983 0.708008ZM11.1266 29.2057H6.22878V13.9774H11.1266V29.2057ZM8.84709 11.662H8.81198C7.24451 11.662 6.22954 10.4952 6.22954 9.03456C6.22954 7.54721 7.27122 6.40785 8.8738 6.40785C10.4764 6.40785 11.457 7.53882 11.4921 9.03456C11.4913 10.4952 10.4764 11.662 8.84709 11.662ZM29.0266 29.2057H24.1288V20.8792C24.1288 18.8843 23.4161 17.5214 21.6441 17.5214C20.2903 17.5214 19.489 18.4371 19.1326 19.3292C18.9991 19.6498 18.9632 20.0863 18.9632 20.5319V29.2057H14.0654V13.9774H18.9632V16.0966C19.676 15.0816 20.7894 13.621 23.3802 13.621C26.5953 13.621 29.0274 15.7402 29.0274 20.3091L29.0266 29.2057Z" fill="#E8E8E8"/>
        </svg>
      </a>
    </div>
</footer>

<style>
    footer#libery-labs-footer {
        
        display: flex;
        position: absolute;
        box-sizing: border-box;
        /* height: calc(var(--vspacing-7) * 0.53125); */
        width: 100vw;
        background: var(--grey-9);
        left: 0;
        bottom: 0;
        z-index: var(--z-index-t-1);
        align-items: center;
        padding: 42px 225px;
        gap: calc(var(--vspacing-6) * 0.84765625);
    }

    #main-icon-wrapper {
        width: calc(var(--vspacing-7) * 0.79296875);
    }

    #footer-libery-labs-star {
        width: calc(var(--vspacing-6) * 0.7154140625);
    }

    #home-links-wrapper #home-links{
      list-style: none;
      padding-inline-start: var(--vspacing-2);
      margin: 0;
    }
    
    #home-links-wrapper #home-links li a {
      color: var(--grey-2) !important;
      text-decoration: none;
    }

    #socials-wrapper {
        display: flex;
        gap: var(--vspacing-2);
    }

    #socials-wrapper svg {
        width: calc(var(--vspacing-3) * 1.5)
    }
</style>
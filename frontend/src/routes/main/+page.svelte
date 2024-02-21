<script>
  import Friends from "./friends.svelte";
  import Posts_page from "./post.svelte";
  import EventSugest from "./eventsugest.svelte";
  import Sugesgroup from "./sugesgroup.svelte";

  import Form from "./createpost.svelte";
  import Comment from "./comments.svelte";
  import Contacts from "./messagerie.svelte";
  import { onMount } from "svelte";

  export let form;
  export let data;
  console.log(data.posts);

  let CommSection = { display: "none", data: {} };

  // console.log("current user", localStorageObj.data);
  let listPostsLength = data?.posts.length;
  let offset = 0;
  $: posts = data?.posts.slice(0, offset);

  let loadingDots;

  let options = {
    rootMargin: "50px",
    threshold: 1.0,
  };

  onMount(() => {
    // window.scrollTo(0, 0);
    let observer = new IntersectionObserver((entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          let timeout = setTimeout(() => {
            offset += 10;
          }, 800);

          //***** this disable the observer when all the posts have been loaded ******/
          if (offset >= listPostsLength) {
            observer.unobserve(loadingDots);
            clearTimeout(timeout);
            loadingDots.style.display = "none";
          }
        }
      });
    }, options);
    observer.observe(loadingDots);
  });
</script>

<!-- main
 content -->
<div class="main-content right-chat-active">
  <div class="middle-sidebar-bottom">
    <!-- <div class="middle-sidebar-left">
      loader wrapper 
      <div class="preloader-wrap p-3">
        <div class="box shimmer">
          <div class="lines">
            <div class="line s_shimmer"></div>
            <div class="line s_shimmer"></div>
            <div class="line s_shimmer"></div>
            <div class="line s_shimmer"></div>
          </div>
        </div>
      </div>
      <div class="box shimmer mb-3">
        <div class="lines">
          <div class="line s_shimmer"></div>
          <div class="line s_shimmer"></div>
          <div class="line s_shimmer"></div>
          <div class="line s_shimmer"></div>
        </div>
      </div>
      <div class="box shimmer">
        <div class="lines">
          <div class="line s_shimmer"></div>
          <div class="line s_shimmer"></div>
          <div class="line s_shimmer"></div>
          <div class="line s_shimmer"></div>
        </div>
      </div>
    </div> -->
    <!-- loader wrapper -->
    <div class="row feed-body">
      <div class="col-xl-8 col-xxl-9 col-lg-8">
        <div class="card w-100 text-center shadow-xss rounded-xxl border-0 p-4 mb-3 mt-3">
          <Form {form} users={data?.folow_and_following} />

          <h1>HERE IS THE PLACE OF POSTS</h1>

          <Posts_page data={posts} bind:CommSection />

          <div bind:this={loadingDots} class="card w-100 text-center shadow-xss rounded-xxl border-0 p-4 mb-3 mt-3">
            <div class="snippet mt-2 ms-auto me-auto" data-title=".dot-typing">
              <div class="stage">
                  <div class="dot-typing"></div>
              </div>
            </div>
          </div>
          
        </div>
      </div>
      <div class="col-xl-4 col-xxl-3 col-lg-4 ps-lg-0">
        <Friends friends={data?.listesusers} />
        <Sugesgroup />
        <EventSugest />
      </div>
    </div>
  </div>
</div>
<!-- main content -->
<Contacts Users={data?.contacts}/>

<Comment bind:CommSection />

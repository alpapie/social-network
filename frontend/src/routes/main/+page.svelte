<script>
  import Friends from "./friends.svelte";
  import Posts from "./posts.svelte";
  import EventSugest from "./eventsugest.svelte";
  import Sugesgroup from "./sugesgroup.svelte";
  import Createpost from "./createpost.svelte";
  export let data ;
</script>
<!-- main content -->
<div class="main-content right-chat-active">
  <div class="middle-sidebar-bottom">
    <div class="middle-sidebar-left">
      <!-- loader wrapper -->
      <div class="preloader-wrap p-3">
        <div class="box shimmer">
          <div class="lines">
            <div class="line s_shimmer"></div>
            <div class="line s_shimmer"></div>
            <div class="line s_shimmer"></div>
            <div class="line s_shimmer"></div>
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
      </div>
      <!-- loader wrapper -->
      <div class="row feed-body">
        <div class="col-xl-8 col-xxl-9 col-lg-8">
          <div
            class="card w-100 text-center shadow-xss rounded-xxl border-0 p-4 mb-3 mt-3"
          >
          <Createpost />
          <Posts posts={data?.posts}/>
            <div class="snippet mt-2 ms-auto me-auto" data-title=".dot-typing">
              <div class="stage">
                <div class="dot-typing"></div>
              </div>
            </div>
          </div>
        </div>
        <div class="col-xl-4 col-xxl-3 col-lg-4 ps-lg-0">
          
         <Friends friends={data?.listesusers}/>
         <Sugesgroup/>
         <EventSugest/>

        </div>
      </div>
    </div>
  </div>
</div>
<!-- main content -->

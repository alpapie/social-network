<script>
    export let CommSection 
</script>

<div id="lightboxOverlay" tabindex="-1" class="lightboxOverlay"></div>
<div id="lightbox" tabindex="-1" class="lightbox">
  
  <div class="lb-dataContainer">
    <div class="lb-data">
      <div class="lb-details">
        <span class="lb-caption"></span><span class="lb-number"></span>
      </div>
     
    </div>
  </div>
  <div class="right-comment chat-left scroll-bar theme-dark-bg" style="display: {CommSection?.display};">
    <div class="card-body ps-2 pe-4 pb-0 d-flex">
      <figure class="avatar me-3">
        <img
          src="images/user-8.png"
          alt="display photo"
          class="shadow-sm rounded-circle w45"
        />
      </figure>
      <h4 class="fw-700 text-grey-900 font-xssss mt-1 text-left">
        {CommSection?.data?.data?.firstName} {CommSection?.data?.data?.lastName} <span
          class="d-block font-xssss fw-500 mt-1 lh-3 text-grey-500"
          >2 hour ago</span
        >
      </h4>
      <a href="/" class="ms-auto"
        ><i class="ti-more-alt text-grey-900 btn-round-md bg-greylight font-xss"
        ></i></a
      >
    </div>
    <div class="card-body d-flex ps-2 pe-4 pt-0 mt-0">
      <a
        href="/"
        class="d-flex align-items-center fw-600 text-grey-900 lh-26 font-xssss me-3 text-dark"
        ><i
          class="feather-thumbs-up text-white bg-primary-gradiant me-1 btn-round-xs font-xss"
        ></i>
        <i
          class="feather-heart text-white bg-red-gradiant me-2 btn-round-xs font-xss"
        ></i>2.8K Like</a
      >
      <a
        href="/"
        class="d-flex align-items-center fw-600 text-grey-900 lh-26 font-xssss text-dark"
        ><i
          class="feather-message-circle text-grey-900 btn-round-sm font-lg text-dark"
        ></i>22 Comment</a
      >
    </div>
    <div class="card w-100 border-0 shadow-none right-scroll-bar">
        {#if CommSection?.data?.data?.Comments && CommSection?.data?.data?.Comments.length > 0}
       
        {#each CommSection?.data?.data?.Comments as com }
            <div class="card-body border-top-xs pt-4 pb-3 pe-4 d-block ps-5">
                <figure class="avatar position-absolute left-0 ms-2 mt-1">
                <img
                    src="images/user-6.png"
                    alt="display photo"
                    class="shadow-sm rounded-circle w35"
                />
                </figure>
                <div
                class="chat p-3 bg-greylight rounded-xxl d-block text-left theme-dark-bg"
                >
                <h4 class="fw-700 text-grey-900 font-xssss mt-0 mb-1">
                    {com.firstName +' '+com.lastName} <a href="/" class="ms-auto"
                    ><i class="ti-more-alt float-right text-grey-800 font-xsss"
                    ></i></a
                    >
                </h4>
                <p class="fw-500 text-grey-500 lh-20 font-xssss w-100 mt-2 mb-0">
                    {com?.comment}
                </p>
                </div>
            </div>
        {/each}
        {/if}
    </div>
  </div>
</div>

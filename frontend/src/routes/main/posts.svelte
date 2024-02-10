
<script>
  export let posts
  function calculateAgo(dateString) {
    // Convert the date string to a Date object
    const date = new Date(dateString);
    const now = new Date();

    // Calculate the difference between the dates in milliseconds
    const difference = now - date;

    // Convert the difference to minutes, hours, or days
    const minutes = Math.floor(difference / (1000 * 60));
    const hours = Math.floor(difference / (1000 * 60 * 60));
    const days = Math.floor(difference / (1000 * 60 * 60 * 24));

    // Determine the age
    if (days > 0) {
        return days +( days>1?" days ago" : " days ago");
    } else if (hours > 0) {
        return hours + (hours>1 ?" hours ago":" hour ago");
    } else {
        return minutes + (minutes>1 ?" minutes ago":" minute ago")
    }
}
  let dd={
    firstName: 'Mamoudou',
    lastName: 'Ndiaye',
    Group_id: 0,
    User_id: 1,
    titre: '2iiiiwie',
    image: 'ada',
    content: 'abbabas',
    privacy: 'public',
    creationDate: '2024-02-09 15:41:09',
    AllowedUsers: null,
    groupName: '',
    description: ''
  }
</script>
{#if posts}
  {#each posts as post}
    <div class="card w-100 shadow-xss rounded-xxl border-0 p-4 mb-4">
      <div class="card-body p-0 d-flex">
        <figure class="avatar me-3">
          {#if !post.image}
              <img src="//ui-avatars.com/api/?name={post.firstName +' '+post.lastName}e&size=100&rounded=true&color=fff&background=random"
            alt="{post.firstName +' '+post.lastName}"
            class="shadow-sm rounded-circle w45"
          />
          {:else}
          <img src="/images/upload/{post.image}" alt="{post.firstName}" class="shadow-sm rounded-circle w45"/>
          {/if}
          
        </figure>
        <h4 class="fw-700 text-grey-900 font-xssss mt-1">
          {post.firstName +' '+post.lastName} <span
            class="d-block font-xssss fw-500 mt-1 lh-3 text-grey-500"
            >{calculateAgo(post.creationDate)}</span
          >
        </h4>
      </div>
      <div class="card-body p-0 me-lg-5">
        <p class="fw-500 text-grey-500 lh-26 font-xssss w-100 mb-2">
          {post.content.slice(0,200)}
          <a
            href="/main/post?post_id={post.id}"
            class="fw-600 text-primary ms-2">See more</a
          >
        </p>
      </div>
      <div class="card-body d-block p-0 mb-3">
        <div class="row ps-2 pe-2">
          <div class="col-sm-12 p-1">
            <a href="/images/t-30.jpg" data-lightbox="roadtr"
              ><img
                src="/images/t-31.jpg"
                class="rounded-3 w-100"
                alt=""
              /></a
            >
          </div>
        </div>
      </div>
      <div class="card-body d-flex p-0">
        <a
          href="/main/post?post_id={post.id}"
          class="d-flex align-items-center fw-600 text-grey-900 text-dark lh-26 font-xssss"
          ><i
            class="feather-message-circle text-dark text-grey-900 btn-round-sm font-lg"
          ></i><span class="d-none-xss">22 Comment</span></a
        >
      </div>
    </div>
  {/each}
{/if}
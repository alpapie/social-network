<script>
  export let friends;
  console.log(friends);
</script>

<div class="card w-100 shadow-xss rounded-xxl border-0 mb-3">
  <div class="card-body d-flex align-items-center p-4">
    <h4 class="fw-700 mb-0 font-xssss text-grey-900">Friend Request</h4>
    <a href="/main/follow" class="fw-600 ms-auto font-xssss text-primary"
      >See all</a
    >
  </div>
  {#if  friends}
    {#each friends as friend}
    <div class="card-body d-flex pt-4 ps-4 pe-4 pb-0 border-top-xs bor-0">
      <figure class="avatar me-3">
        {#if !friend.Avatar }
        <img src="{friend.Avatar}" alt="{friend.FirstName +" "+ friend.LastName}" class="float-right p-0 bg-white rounded-circle w-100 shadow-xss">
        {:else}
        <img src="//ui-avatars.com/api/?name={friend.FirstName +" "+ friend.LastName}e&size=100&rounded=true&color=fff&background=random" alt="{user.FirstName +" "+ user.LastName}" />

        {/if}
        
      </figure>
      <h4 class="fw-700 text-grey-900 font-xssss mt-1">
        {friend.FirstName+" "+
        friend.LastName} 
        <!-- <span
        class="d-block font-xssss fw-500 mt-1 lh-3 text-grey-500"
        >12 mutual friends</span -->
        <!-- > -->
      </h4>
    </div>
    <div class="card-body d-flex align-items-center pt-0 ps-4 pe-4 pb-4">
      <a href="#" class="mt-0 btn pt-2 pb-2 ps-3 pe-3 lh-24 ms-1 ls-3 d-inline-block rounded-xl bg-success font-xsssss fw-700 ls-lg text-white">ADD FRIEND</a>
    </div>
    {/each}
  {/if}

  
</div>

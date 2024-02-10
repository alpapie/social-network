<script>
	export let data;
</script>

{#each data as d}
<div class="card w-100 shadow-xss rounded-xxl border-0 p-4 mb-3">
	{#if d.Group_id != 0}
	<div class="card-body p-0 d-flex">
		<figure class="avatar me-3">
			<img src="images/user-7.png" alt="post" class="shadow-sm rounded-circle w45" />
		</figure>
		<h4 class="fw-500 text-grey-900 font-xs mt-1">
			{d.groupName}
			<!-- <span class="d-block font-xssss fw-500 mt-1 lh-3 text-grey-500">3 hour ago</span> -->
		</h4>
	</div>
	{/if}
	<div class="card-body p-0 d-flex">
		<figure class="avatar me-3">
			<img src="images/user-7.png" alt="post" class="shadow-sm rounded-circle w45" />
		</figure>
		<h4 class="fw-700 text-grey-900 font-xssss mt-1">
			{d.firstName}
			{d.lastName}
			<span class="d-block font-xssss fw-500 mt-1 lh-3 text-grey-500">3 hour ago</span>
		</h4>
	</div>
	<div class="card-body p-0 me-lg-5">
		<p class="fw-500 text-grey-500 lh-26 font-xssss w-100">
			{d.titre}
			{d.content}
			{#if d.content.length > 300}
				<a href="/" class="fw-600 text-primary ms-2">See more</a>
			{/if}
		</p>
	</div>
{#if d.image != ""}	
	<div class="card-body d-block p-0">
		<div class="row ps-2 pe-2">
			<div class="col-xs-4 col-sm-4 p-1" style="width: 30rem;">
				<a href="images/t-10.jpg" data-lightbox="roadtrip"><img src={d.image} class="rounded-3 w-100"
						alt="random" /></a>
			</div>
		</div>
	</div>
{/if}
	<div class="card-body d-flex p-0 mt-3">
		<a href="/postdetail/?postid={d.id}" class="d-flex align-items-center fw-600 text-grey-900 text-dark lh-26 font-xssss"><i
				class="feather-message-circle text-dark text-grey-900 btn-round-sm font-lg"></i><span
				class="d-none-xss">22 Comment</span></a>
	</div>
</div>
{/each}
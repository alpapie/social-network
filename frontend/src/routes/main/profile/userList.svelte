<script>
	export let users;
	export let isfollowing;
	export let unfollow;
</script>

<div class="row">
	<div class="col-xl-12">
		<div class="card shadow-xss w-100 d-block d-flex border-0 p-4 mb-3">
			<div class="card-body d-flex align-items-center p-0">
				<h2 class="fw-700 mb-0 mt-0 font-md text-grey-900">Friends</h2>
				<div class="search-form-2 ms-auto">
					<i class="ti-search font-xss"></i>
					<input
						type="text"
						class="form-control text-grey-500 mb-0 bg-greylight theme-dark-bg border-0"
						placeholder="Search here."
					/>
				</div>
				<a
					href="#"
					class="btn-round-md ms-2 bg-greylight theme-dark-bg rounded-3"
					><i class="feather-filter font-xss text-grey-500"></i></a
				>
			</div>
		</div>

		<div class="row ps-2 pe-2">
			{#each users as user}
				<div class="col-md-3 col-sm-4 pe-2 ps-2">
					<div
						class="card d-block border-0 shadow-xss rounded-3 overflow-hidden mb-3"
					>
						<div class="card-body d-block w-100 ps-3 pe-3 pb-4 text-center">
							<figure
								class="avatar ms-auto me-auto mb-0 position-relative w100 z-index-1"
							>
								{#if user.Avatar}
									<img
										src={user.Avatar}
										alt={user.FirstName + " " + user.LastName}
										class="float-right p-0 bg-white rounded-circle w-100 shadow-xss"
									/>
								{:else}
									<img
										src="//ui-avatars.com/api/?name={user.FirstName +
											' ' +
											user.LastName}e&size=100&rounded=true&color=fff&background=random"
										alt={user.FirstName + " " + user.LastName}
									/>
								{/if}
							</figure>

							{#if isfollowing}
								<a href="/main/profile/{user.ID}">
									<h4 class="fw-700 font-xsss mt-3 mb-1">
										{user.FirstName + " " + user.LastName}
									</h4>
								</a>
							{:else if user.IsPublic === 1}
								<a href="/main/profile/{user.ID}">
									<h4 class="fw-700 font-xsss mt-3 mb-1">
										{user.FirstName + " " + user.LastName}
									</h4>
								</a>
							{:else}
								<h4 class="fw-700 font-xsss mt-3 mb-1">
									{user.FirstName + " " + user.LastName}
								</h4>
							{/if}
							<p class="fw-500 font-xsssss text-grey-500 mt-0 mb-3">
								{user.Email}
							</p>
							{#if isfollowing}
								<button
									on:click={unfollow(user.ID)}
									class="mt-0 btn pt-2 pb-2 ps-3 pe-3 lh-24 b-0 ms-1 ls-3 d-inline-block rounded-xl bg-danger font-xsssss fw-700 ls-lg text-white"
									>UNFOLLOW
								</button>
							{/if}
						</div>
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>

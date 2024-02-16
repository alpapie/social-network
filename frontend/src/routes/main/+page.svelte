<script>
	import Form from "./form.svelte";
	import Post from "./post.svelte";
	import Comment from "./comments.svelte";
	import { onMount } from "svelte";
	export let data;
	export let form;
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
		<div class="middle-sidebar-left">
			<!-- loader wrapper -->
			<!-- <div class="preloader-wrap p-3">c
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
        </div> -->
			<!-- loader wrapper -->
			<div class="row feed-body" style="overflow: auto; height: auto">
				<div class="col-xl-8 col-xxl-9 col-lg-8">
					<div
						class="card w-100 shadow-none bg-transparent bg-transparent-card border-0 p-0 mb-0"
					>
						<div
							class="owl-carousel category-card owl-theme overflow-hidden nav-none"
						>
							<div class="item">
								<div
									data-bs-toggle="modal"
									data-bs-target="#Modalstory"
									class="card w125 h200 d-block border-0 shadow-none rounded-xxxl bg-dark overflow-hidden mb-3 mt-3"
								>
									<div
										class="card-body d-block p-3 w-100 position-absolute bottom-0 text-center"
									>
										<a href="/">
											<span class="btn-round-lg bg-white"
												><i class="feather-plus font-lg"></i></span
											>
											<div class="clearfix"></div>
											<h4
												class="fw-700 position-relative z-index-1 ls-1 font-xssss text-white mt-2 mb-1"
											>
												Add Story
											</h4>
										</a>
									</div>
								</div>
							</div>
							<div class="item">
								<div
									data-bs-toggle="modal"
									data-bs-target="#Modalstory"
									class="card w125 h200 d-block border-0 shadow-xss rounded-xxxl bg-gradiant-bottom overflow-hidden cursor-pointer mb-3 mt-3"
									style="background-image: url(images/s-1.jpg);"
								>
									<div
										class="card-body d-block p-3 w-100 position-absolute bottom-0 text-center"
									>
										<a href="/">
											<figure
												class="avatar ms-auto me-auto mb-0 position-relative w50 z-index-1"
											>
												<img
													src="images/user-11.png"
													alt="random"
													class="float-right p-0 bg-white rounded-circle w-100 shadow-xss"
												/>
											</figure>
											<div class="clearfix"></div>
											<h4
												class="fw-600 position-relative z-index-1 ls-1 font-xssss text-white mt-2 mb-1"
											>
												Victor Exrixon
											</h4>
										</a>
									</div>
								</div>
							</div>
							<div class="item">
								<div
									data-bs-toggle="modal"
									data-bs-target="#Modalstory"
									class="card w125 h200 d-block border-0 shadow-xss rounded-xxxl bg-gradiant-bottom overflow-hidden cursor-pointer mb-3 mt-3"
									style="background-image: url(images/s-2.jpg);"
								>
									<div
										class="card-body d-block p-3 w-100 position-absolute bottom-0 text-center"
									>
										<a href="/">
											<figure
												class="avatar ms-auto me-auto mb-0 position-relative w50 z-index-1"
											>
												<img
													src="images/user-12.png"
													alt="random"
													class="float-right p-0 bg-white rounded-circle w-100 shadow-xss"
												/>
											</figure>
											<div class="clearfix"></div>
											<h4
												class="fw-600 position-relative z-index-1 ls-1 font-xssss text-white mt-2 mb-1"
											>
												Surfiya Zakir
											</h4>
										</a>
									</div>
								</div>
							</div>
							<div class="item">
								<div
									data-bs-toggle="modal"
									data-bs-target="#Modalstory"
									class="card w125 h200 d-block border-0 shadow-xss rounded-xxxl bg-gradiant-bottom overflow-hidden cursor-pointer mb-3 mt-3"
								>
									<!-- <video autoplay loop class="float-right w-100">
                    <source src="images/s-4.mp4" type="video/mp4" />
                  </video> -->
									<div
										class="card-body d-block p-3 w-100 position-absolute bottom-0 text-center"
									>
										<a href="/">
											<figure
												class="avatar ms-auto me-auto mb-0 position-relative w50 z-index-1"
											>
												<img
													src="images/user-9.png"
													alt="random"
													class="float-right p-0 bg-white rounded-circle w-100 shadow-xss"
												/>
											</figure>
											<div class="clearfix"></div>
											<h4
												class="fw-600 position-relative z-index-1 ls-1 font-xssss text-white mt-2 mb-1"
											>
												Goria Coast
											</h4>
										</a>
									</div>
								</div>
							</div>
							<div class="item">
								<div
									data-bs-toggle="modal"
									data-bs-target="#Modalstory"
									class="card w125 h200 d-block border-0 shadow-xss rounded-xxxl bg-gradiant-bottom overflow-hidden cursor-pointer mb-3 mt-3 me-1"
								>
									<!-- <video autoplay loop class="float-right w-100">
                    <source src="images/s-3.mp4" type="video/mp4" />
                  </video> -->
									<div
										class="card-body d-block p-3 w-100 position-absolute bottom-0 text-center"
									>
										<a href="/">
											<figure
												class="avatar ms-auto me-auto mb-0 position-relative w50 z-index-1"
											>
												<img
													src="images/user-4.png"
													alt="random"
													class="float-right p-0 bg-white rounded-circle w-100 shadow-xss"
												/>
											</figure>
											<div class="clearfix"></div>
											<h4
												class="fw-600 position-relative z-index-1 ls-1 font-xssss text-white mt-2 mb-1"
											>
												Hurin Seary
											</h4>
										</a>
									</div>
								</div>
							</div>
							<div class="item">
								<div
									data-bs-toggle="modal"
									data-bs-target="#Modalstory"
									class="card w125 h200 d-block border-0 shadow-xss rounded-xxxl bg-gradiant-bottom overflow-hidden cursor-pointer mb-3 mt-3"
									style="background-image: url(images/s-5.jpg);"
								>
									<div
										class="card-body d-block p-3 w-100 position-absolute bottom-0 text-center"
									>
										<a href="/">
											<figure
												class="avatar ms-auto me-auto mb-0 position-relative w50 z-index-1"
											>
												<img
													src="images/user-3.png"
													alt="random"
													class="float-right p-0 bg-white rounded-circle w-100 shadow-xss"
												/>
											</figure>
											<div class="clearfix"></div>
											<h4
												class="fw-600 position-relative z-index-1 ls-1 font-xssss text-white mt-2 mb-1"
											>
												David Goria
											</h4>
										</a>
									</div>
								</div>
							</div>
							<div class="item">
								<div
									data-bs-toggle="modal"
									data-bs-target="#Modalstory"
									class="card w125 h200 d-block border-0 shadow-xss rounded-xxxl bg-gradiant-bottom overflow-hidden cursor-pointer mb-3 mt-3"
									style="background-image: url(images/s-6.jpg);"
								>
									<div
										class="card-body d-block p-3 w-100 position-absolute bottom-0 text-center"
									>
										<a href="/">
											<figure
												class="avatar ms-auto me-auto mb-0 position-relative w50 z-index-1"
											>
												<img
													src="images/user-2.png"
													alt="random"
													class="float-right p-0 bg-white rounded-circle w-100 shadow-xss"
												/>
											</figure>
											<div class="clearfix"></div>
											<h4
												class="fw-600 position-relative z-index-1 ls-1 font-xssss text-white mt-2 mb-1"
											>
												Seary Victor
											</h4>
										</a>
									</div>
								</div>
							</div>
						</div>
					</div>

					<Form {form} users={data?.listesusers} />

					<h1>HERE IS THE PLACE OF POSTS</h1>

					<Post data={posts} bind:CommSection />

					<div
						bind:this={loadingDots}
						class="card w-100 text-center shadow-xss rounded-xxl border-0 p-4 mb-3 mt-3"
					>
						<div class="snippet mt-2 ms-auto me-auto" data-title=".dot-typing">
							<div class="stage">
								<div class="dot-typing"></div>
							</div>
						</div>
					</div>
				</div>
				<div class="col-xl-4 col-xxl-3 col-lg-4 ps-lg-0">
					<div class="card w-100 shadow-xss rounded-xxl border-0 mb-3">
						<div class="card-body d-flex align-items-center p-4">
							<h4 class="fw-700 mb-0 font-xssss text-grey-900">
								Friend Request
							</h4>
							<a
								href="default-member.html"
								class="fw-600 ms-auto font-xssss text-primary">See all</a
							>
						</div>
						<div
							class="card-body d-flex pt-4 ps-4 pe-4 pb-0 border-top-xs bor-0"
						>
							<figure class="avatar me-3">
								<img
									src="images/user-7.png"
									alt="random"
									class="shadow-sm rounded-circle w45"
								/>
							</figure>
							<h4 class="fw-700 text-grey-900 font-xssss mt-1">
								Anthony Daugloi <span
									class="d-block font-xssss fw-500 mt-1 lh-3 text-grey-500"
									>12 mutual friends</span
								>
							</h4>
						</div>
						<div
							class="card-body d-flex align-items-center pt-0 ps-4 pe-4 pb-4"
						>
							<a
								href="/"
								class="p-2 lh-20 w100 bg-primary-gradiant me-2 text-white text-center font-xssss fw-600 ls-1 rounded-xl"
								>Confirm</a
							>
							<a
								href="/"
								class="p-2 lh-20 w100 bg-grey text-grey-800 text-center font-xssss fw-600 ls-1 rounded-xl"
								>Delete</a
							>
						</div>

						<div class="card-body d-flex pt-0 ps-4 pe-4 pb-0">
							<figure class="avatar me-3">
								<img
									src="images/user-8.png"
									alt="random"
									class="shadow-sm rounded-circle w45"
								/>
							</figure>
							<h4 class="fw-700 text-grey-900 font-xssss mt-1">
								Mohannad Zitoun <span
									class="d-block font-xssss fw-500 mt-1 lh-3 text-grey-500"
									>12 mutual friends</span
								>
							</h4>
						</div>
						<div
							class="card-body d-flex align-items-center pt-0 ps-4 pe-4 pb-4"
						>
							<a
								href="/"
								class="p-2 lh-20 w100 bg-primary-gradiant me-2 text-white text-center font-xssss fw-600 ls-1 rounded-xl"
								>Confirm</a
							>
							<a
								href="/"
								class="p-2 lh-20 w100 bg-grey text-grey-800 text-center font-xssss fw-600 ls-1 rounded-xl"
								>Delete</a
							>
						</div>

						<div class="card-body d-flex pt-0 ps-4 pe-4 pb-0">
							<figure class="avatar me-3">
								<img
									src="images/user-12.png"
									alt="random"
									class="shadow-sm rounded-circle w45"
								/>
							</figure>
							<h4 class="fw-700 text-grey-900 font-xssss mt-1">
								Mohannad Zitoun <span
									class="d-block font-xssss fw-500 mt-1 lh-3 text-grey-500"
									>12 mutual friends</span
								>
							</h4>
						</div>
						<div
							class="card-body d-flex align-items-center pt-0 ps-4 pe-4 pb-4"
						>
							<a
								href="/"
								class="p-2 lh-20 w100 bg-primary-gradiant me-2 text-white text-center font-xssss fw-600 ls-1 rounded-xl"
								>Confirm</a
							>
							<a
								href="/"
								class="p-2 lh-20 w100 bg-grey text-grey-800 text-center font-xssss fw-600 ls-1 rounded-xl"
								>Delete</a
							>
						</div>
					</div>

					<div class="card w-100 shadow-xss rounded-xxl border-0 p-0">
						<div class="card-body d-flex align-items-center p-4 mb-0">
							<h4 class="fw-700 mb-0 font-xssss text-grey-900">
								Confirm Friend
							</h4>
							<a
								href="default-member.html"
								class="fw-600 ms-auto font-xssss text-primary">See all</a
							>
						</div>
						<div
							class="card-body bg-transparent-card d-flex p-3 bg-greylight ms-3 me-3 rounded-3"
						>
							<figure class="avatar me-2 mb-0">
								<img
									src="images/user-7.png"
									alt="random"
									class="shadow-sm rounded-circle w45"
								/>
							</figure>
							<h4 class="fw-700 text-grey-900 font-xssss mt-2">
								Anthony Daugloi <span
									class="d-block font-xssss fw-500 mt-1 lh-3 text-grey-500"
									>12 mutual friends</span
								>
							</h4>
							<!-- <a
                href="/"
                class="btn-round-sm bg-white text-grey-900 feather-chevron-right font-xss ms-auto mt-2"
              ></a> -->
						</div>
						<div
							class="card-body bg-transparent-card d-flex p-3 bg-greylight m-3 rounded-3"
							style="margin-bottom: 0 !important;"
						>
							<figure class="avatar me-2 mb-0">
								<img
									src="images/user-8.png"
									alt="random"
									class="shadow-sm rounded-circle w45"
								/>
							</figure>
							<h4 class="fw-700 text-grey-900 font-xssss mt-2">
								David Agfree <span
									class="d-block font-xssss fw-500 mt-1 lh-3 text-grey-500"
									>12 mutual friends</span
								>
							</h4>
							<!-- <a
                href="/"
                class="btn-round-sm bg-white text-grey-900 feather-plus font-xss ms-auto mt-2"
              ></a> -->
						</div>
						<div
							class="card-body bg-transparent-card d-flex p-3 bg-greylight m-3 rounded-3"
						>
							<figure class="avatar me-2 mb-0">
								<img
									src="images/user-12.png"
									alt="random"
									class="shadow-sm rounded-circle w45"
								/>
							</figure>
							<h4 class="fw-700 text-grey-900 font-xssss mt-2">
								Hugury Daugloi <span
									class="d-block font-xssss fw-500 mt-1 lh-3 text-grey-500"
									>12 mutual friends</span
								>
							</h4>
							<!-- <a
                href="/"
                class="btn-round-sm bg-white text-grey-900 feather-plus font-xss ms-auto mt-2"
              ></a> -->
						</div>
					</div>

					<div class="card w-100 shadow-xss rounded-xxl border-0 mb-3 mt-3">
						<div class="card-body d-flex align-items-center p-4">
							<h4 class="fw-700 mb-0 font-xssss text-grey-900">
								Suggest Group
							</h4>
							<a
								href="default-group.html"
								class="fw-600 ms-auto font-xssss text-primary">See all</a
							>
						</div>
						<div
							class="card-body d-flex pt-4 ps-4 pe-4 pb-0 overflow-hidden border-top-xs bor-0"
						>
							<img
								src="images/e-2.jpg"
								alt="img"
								class="img-fluid rounded-xxl mb-2"
							/>
						</div>
						<div class="card-body dd-block pt-0 ps-4 pe-4 pb-4">
							<ul class="memberlist mt-1 mb-2 ms-0 d-block">
								<li class="w20">
									<a href="/"
										><img
											src="images/user-6.png"
											alt="user"
											class="w35 d-inline-block"
											style="opacity: 1;"
										/></a
									>
								</li>
								<li class="w20">
									<a href="/"
										><img
											src="images/user-7.png"
											alt="user"
											class="w35 d-inline-block"
											style="opacity: 1;"
										/></a
									>
								</li>
								<li class="w20">
									<a href="/"
										><img
											src="images/user-8.png"
											alt="user"
											class="w35 d-inline-block"
											style="opacity: 1;"
										/></a
									>
								</li>
								<li class="w20">
									<a href="/"
										><img
											src="images/user-3.png"
											alt="user"
											class="w35 d-inline-block"
											style="opacity: 1;"
										/></a
									>
								</li>
								<li class="last-member">
									<a
										href="/"
										class="bg-greylight fw-600 text-grey-500 font-xssss w35 ls-3 text-center"
										style="height: 35px; line-height: 35px;">+2</a
									>
								</li>
								<li class="ps-3 w-auto ms-1">
									<a href="/" class="fw-600 text-grey-500 font-xssss"
										>Member apply</a
									>
								</li>
							</ul>
						</div>
					</div>

					<div class="card w-100 shadow-xss rounded-xxl border-0 mb-3">
						<div class="card-body d-flex align-items-center p-4">
							<h4 class="fw-700 mb-0 font-xssss text-grey-900">
								Suggest Pages
							</h4>
							<a
								href="default-group.html"
								class="fw-600 ms-auto font-xssss text-primary">See all</a
							>
						</div>
						<div
							class="card-body d-flex pt-4 ps-4 pe-4 pb-0 overflow-hidden border-top-xs bor-0"
						>
							<img
								src="images/g-2.jpg"
								alt="img"
								class="img-fluid rounded-xxl mb-2"
							/>
						</div>
						<div
							class="card-body d-flex align-items-center pt-0 ps-4 pe-4 pb-4"
						>
							<a
								href="/"
								class="p-2 lh-28 w-100 bg-grey text-grey-800 text-center font-xssss fw-700 rounded-xl"
								><i class="feather-external-link font-xss me-2"></i> Like Page</a
							>
						</div>

						<div class="card-body d-flex pt-0 ps-4 pe-4 pb-0 overflow-hidden">
							<img
								src="images/g-3.jpg"
								alt="img"
								class="img-fluid rounded-xxl mb-2 bg-lightblue"
							/>
						</div>
						<div
							class="card-body d-flex align-items-center pt-0 ps-4 pe-4 pb-4"
						>
							<a
								href="/"
								class="p-2 lh-28 w-100 bg-grey text-grey-800 text-center font-xssss fw-700 rounded-xl"
								><i class="feather-external-link font-xss me-2"></i> Like Page</a
							>
						</div>
					</div>

					<div class="card w-100 shadow-xss rounded-xxl border-0 mb-3">
						<div class="card-body d-flex align-items-center p-4">
							<h4 class="fw-700 mb-0 font-xssss text-grey-900">Event</h4>
							<a
								href="default-event.html"
								class="fw-600 ms-auto font-xssss text-primary">See all</a
							>
						</div>
						<div class="card-body d-flex pt-0 ps-4 pe-4 pb-3 overflow-hidden">
							<div class="bg-success me-2 p-3 rounded-xxl">
								<h4 class="fw-700 font-lg ls-3 lh-1 text-white mb-0">
									<span class="ls-1 d-block font-xsss text-white fw-600"
										>FEB</span
									>22
								</h4>
							</div>
							<h4 class="fw-700 text-grey-900 font-xssss mt-2">
								Meeting with clients <span
									class="d-block font-xsssss fw-500 mt-1 lh-4 text-grey-500"
									>41 madison ave, floor 24 new work, NY 10010</span
								>
							</h4>
						</div>

						<div class="card-body d-flex pt-0 ps-4 pe-4 pb-3 overflow-hidden">
							<div class="bg-warning me-2 p-3 rounded-xxl">
								<h4 class="fw-700 font-lg ls-3 lh-1 text-white mb-0">
									<span class="ls-1 d-block font-xsss text-white fw-600"
										>APR</span
									>30
								</h4>
							</div>
							<h4 class="fw-700 text-grey-900 font-xssss mt-2">
								Developer Programe <span
									class="d-block font-xsssss fw-500 mt-1 lh-4 text-grey-500"
									>41 madison ave, floor 24 new work, NY 10010</span
								>
							</h4>
						</div>

						<div class="card-body d-flex pt-0 ps-4 pe-4 pb-3 overflow-hidden">
							<div class="bg-primary me-2 p-3 rounded-xxl">
								<h4 class="fw-700 font-lg ls-3 lh-1 text-white mb-0">
									<span class="ls-1 d-block font-xsss text-white fw-600"
										>APR</span
									>23
								</h4>
							</div>
							<h4 class="fw-700 text-grey-900 font-xssss mt-2">
								Aniversary Event <span
									class="d-block font-xsssss fw-500 mt-1 lh-4 text-grey-500"
									>41 madison ave, floor 24 new work, NY 10010</span
								>
							</h4>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
<!-- main content -->
<Comment bind:CommSection />

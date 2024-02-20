<script>
	export let CommSection;
	import { enhance } from "$app/forms";
	let inputField;

	const onInput = (event) => {
		if (event.key !== "Enter") return;
		inputField = "";
		console.log("comment content input field",inputField);
	};

	function CloseComment() {
		CommSection.display = "none"
	}
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
		<div
			class="right-comment chat-left scroll-bar theme-dark-bg"
			style="display: {CommSection?.display};"
		>
			<div class="card-body ps-2 pe-4 pb-0 d-flex">
				<figure class="avatar me-3">
					<img
						src="/images/user-8.png"
						alt="display photo"
						class="shadow-sm rounded-circle w45"
					/>
				</figure>
				<h4 class="fw-700 text-grey-900 font-xssss mt-1 text-left">
					{CommSection?.data?.data?.firstName}
					{CommSection?.data?.data?.lastName}
					<span class="d-block font-xssss fw-500 mt-1 lh-3 text-grey-500"
						>2 hour ago</span
					>
				</h4>
				<a href="/" class="ms-auto"  on:click={CloseComment}
					><i class="ti-more-alt text-grey-900 btn-round-md bg-greylight font-xss"
					></i></a
				>
			</div>
			<!-- input -->
			<form
				action="?/createComment"
				method="post"
				class="d-flex"
				enctype="multipart/form-data"
				use:enhance={({ formData }) => {
					formData.append("postId", CommSection?.data?.data?.id);
					return async ({ result }) => {
						
						if (result?.data?.succes == true) {
							if (
								CommSection?.data?.data?.Comments &&
								CommSection?.data?.data?.Comments.length > 0
							) {
								CommSection.data.data.Comments = [
									result.data.data,
									...CommSection.data.data.Comments,
								];
							} else {
								CommSection.data.data.Comments = [result.data.data];
							}
						}
						// console.log("new POSTDETAIL", CommSection.data);
						// CommSection?.data?.data?.Comment = [result?.data?.success , ...CommSection?.data?.data?.Comment]
					};
				}}
			> 
				<button
					class="font-xssss fw-600 text-grey-700 card-body p-0 d-flex align-items-center"
					style="border: none; background-color: inherit;"
					><i
						class="btn-round-sm font-xs text-primary feather-edit-3 me-2 bg-greylight"
					></i>
				</button>
				<!-- <div class="form-group mb-0 icon-input"> -->
				<input
					type="text"
					placeholder="Add new Comment.."
					name="comment"
					class="bg-grey border-0 lh-32 pt-1 pb-1 ps-3 pe-3 font-xssss fw-400 rounded-xl w250 theme-dark-bg"
				/>
				<!-- <input name="postId" type="text" value={CommSection?.data?.data?.id} /> -->
			<div>
					<label style="cursor: pointer;" for="commimage" class="d-flex align-items-center font-xssss fw-600 ls-1 text-grey-700 text-dark pe-4">
						<i class="font-md text-success feather-image me-2"></i>
						<span class="d-none-xs" style="white-space: nowrap;">Insert an image</span>
				   </label>
					<input
					name="avatar"
					id="commimage"
					style="display: none;"
					accept=".jpg, .jpeg, .png, .gif"
					type="file"
					/>
			</div>


			</form>
			<!--  -->
			<div class="card-body d-flex ps-2 pe-4 pt-0 mt-0">
				<a
					href="/"
					class="d-flex align-items-center fw-600 text-grey-900 lh-26 font-xssss text-dark"
					><i
						class="feather-message-circle text-grey-900 btn-round-sm font-lg text-dark"
					></i>{#if CommSection?.data?.data?.Comments}
						{CommSection?.data?.data?.Comments.length}
					{:else}
						0
					{/if}</a
				>
			</div>
			<div class="card w-100 border-0 shadow-none right-scroll-bar">
				{#if CommSection?.data?.data?.Comments && CommSection?.data?.data?.Comments.length > 0}
					{#each CommSection?.data?.data?.Comments as com}
						<div class="card-body border-top-xs pt-4 pb-3 pe-4 d-block ps-5">
							<figure class="avatar position-absolute left-0 ms-2 mt-1">
								<img
									src="/images/user-6.png"
									alt="display photo"
									class="shadow-sm rounded-circle w35"
								/>
							</figure>
							<div
								class="chat p-3 bg-greylight rounded-xxl d-block text-left theme-dark-bg"
							>
								<h4 class="fw-700 text-grey-900 font-xssss mt-0 mb-1">
									{com?.firstName + " " + com?.lastName}
									<a href="/" class="ms-auto"
										><i class="ti-more-alt float-right text-grey-800 font-xsss"
										></i></a
									>
								</h4>
								<p class="fw-500 text-grey-500 lh-20 font-xssss w-100 mt-2 mb-0">
									{com?.comment}
								</p>
								<div class="col-sm-12 p-1" >
									<img 
									  src={com?.image} 
									  class="rounded-3 w-100"
									  alt=""
									/>
								</div>
							</div>
						</div>
					{/each}
				{/if}
			</div>
		</div>
	</div>

<script>
	import { enhance } from "$app/forms";
	export let form;
	export let modalDisplay = "none";
	function handleSelect(e) {
		let selectValue = e.target.value;
		if (selectValue === "almostprivate") {
			modalDisplay = "block";
		} else {
			modalDisplay = "none";
		}
	}
</script>

<form
	method="POST"
	action="?/createPost"
	class="card w-100 shadow-xss rounded-xxl border-0 ps-4 pt-4 pe-4 pb-3 mb-3"
	use:enhance
	enctype="multipart/form-data"
>
	<div
		style="display: flex;justify-content: space-between;"
		class="card-body p-0"
	>
		<div>
			<button
				style="border: none;background-color: inherit;"
				class=" font-xssss fw-600 text-grey-700 card-body p-0 d-flex align-items-center"
			>
				<i
					class="btn-round-sm font-xs text-primary feather-edit-3 me-2 bg-greylight"
				></i>Create Post
			</button>
		</div>
	</div>
	<div class="card-body p-0 mt-3 position-relative">
		<figure class="avatar position-absolute ms-2 mt-1 top-5">
			<img
				src="/images/user-8.png"
				alt="random"
				class="shadow-sm rounded-circle w30"
			/>
		</figure>
		<textarea
			name="content"
			class="h100 bor-0 w-100 rounded-xxl p-2 ps-5 font-xssss text-grey-500 fw-500 border-light-md theme-dark-bg"
			cols="30"
			rows="10"
			placeholder="What's on your mind?"
		></textarea>
		{#if form?.missing}
			<span class="font-xssss" style="color: red;">Empty post</span>
		{/if}
		<div class="card-body d-flex p-0 mt-0">
			<label
				style="cursor: pointer;"
				for="my-image"
				class="d-flex align-items-center font-xssss fw-600 ls-1 text-grey-700 text-dark pe-4"
			>
				<i class="font-md text-success feather-image me-2"></i>
				<span class="d-none-xs" style="white-space: nowrap;"
					>Insert an image</span
				>
			</label>
			<input
				name="avatar"
				id="my-image"
				style="display: none;"
				accept=".jpg, .jpeg, .png, .gif"
				type="file"
			/>
		</div>
	</div>
</form>

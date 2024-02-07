<script>
	import Header from "./header.svelte";
	import SideBar from "./sidebar.svelte";
</script>

<div class="main-wrapper">
	<Header />
	<SideBar />
	<slot />
</div>

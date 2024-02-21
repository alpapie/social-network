<script>
	let activateSide = false
	import Header from "./header.svelte";
	import SideBar from "./sidebar.svelte";
</script>

<div class="main-wrapper">
	<Header {activateSide}/>
	<SideBar />
	<slot />
</div>

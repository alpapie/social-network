<script>
	import Header from "./header.svelte";
	import SideBar from "./sidebar.svelte";
	export let data

</script>

<div class="main-wrapper">
	<Header data={data}/>
	<SideBar />
	<slot />
</div>

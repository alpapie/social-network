<script>
	let activateSide = false
	import Header from "./header.svelte";
	import SideBar from "./sidebar.svelte";
	import Contacts from "./messagerie.svelte";
	let Users = [];
	
</script>

<div class="main-wrapper">
	<Header {activateSide}/>
	<SideBar />
	<slot />
	<Contacts Users={Users}/>
</div>

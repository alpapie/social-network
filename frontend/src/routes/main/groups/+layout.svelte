<script>
    import Header from '../main/header.svelte'
    import SideBar from '../main/sidebar.svelte'
</script>
<div class="main-wrapper">
    <Header/> 
    <SideBar/>
    <slot />
</div>
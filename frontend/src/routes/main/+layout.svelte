<script>
	export let data
	import { PUBLIC_BACKEND_URL } from '$env/static/public';
	let activateSide = false
	import Header from "./header.svelte";
	import SideBar from "./sidebar.svelte";
	import Contacts from "./messagerie.svelte";
	import { onMount } from "svelte";
	import { ContactsStore } from "./stores";
	import axios from "axios";
	let Users = [];
	onMount(async () => { 
		console.log('Mounted');
		let url = `${PUBLIC_BACKEND_URL}/getcontacts`;
		try {
			let header={
					cookie:document.cookie
				}
				const config = { method:"get",withCredentials: true , header,mode: 'no-cors' };
				let response= await axios(url,config)
			    console.log("SUCCES CONTACTS",  response.data.contacts);
				Users = response?.data?.contacts
				ContactsStore.set(Users)
            // return response .data.contacts
        } catch (err) {
          throw console.log("ERORR CONtacts",  err);
        }
      })
</script>

<div class="main-wrapper">
	<Header data={data.notif}  {activateSide}/>
	<SideBar data={data.user} />
	<slot />
	<Contacts />
</div>

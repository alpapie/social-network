<!-- right chat -->
<script>
	import { onDestroy } from "svelte";
	import { displayContacts } from "./stores";
	import { onMount } from "svelte";
	import AlertContainer from "./alert.svelte";
	import { LastMessage } from "./stores";
	import { ContactsStore } from "./stores";
	import { OnlineStore } from "./stores";
	import { OfflineStore } from "./stores";
	import { get } from "svelte/store";

	let show;
	let Users;
	let lastMesage = "";
	let showAlert = false;

	ContactsStore.subscribe((val) => (Users = val));

	const unsuscribe_displayContacts = displayContacts.subscribe((value) => {
		show = value;
	});

	const toggleConstants = () => {
		displayContacts.update((show) => !show);
	};
	import { WS } from "./socket";
	let socket;
	$: usersOnline = [];
	// $:online = []
	// $:offline = []

	onMount(async () => {
		if (
			!window.location.href.includes("main/chat") &&
			!window.location.href.includes("/messages")
		) {
			WS.subscribe((val) => (socket = val));
			socket.addEventListener("open", () => {
				console.log("Opened");
			});
			socket.onmessage = (e) => {
				var newMessage = JSON.parse(e.data);
				if (newMessage.action == "onlineUsers") {
					usersOnline = newMessage.users;
					if (Users && Users.length > 0) {
						for (let user of Users) {
							if (usersOnline.some((onlineUser) => onlineUser.ID === user.ID)) {
								user.Online = true;
							} else {
								user.Online = false;
							}
						}
						ContactsStore.set(Users);
					}
				} else {
					lastMesage = newMessage;
					showAlert = true;
					setTimeout(() => (showAlert = false), 5000);
				}
			};
		}
	});

	let lastmess = undefined;
	const unsuscribeLastMessage = LastMessage.subscribe((val) => {
		lastmess = val;
	});
	$: {
		if (lastmess.content !== undefined) {
			lastMesage = lastmess;
			showAlert = true;
			setTimeout(() => (showAlert = false), 5000);
		}
	}
	onDestroy(unsuscribeLastMessage, unsuscribe_displayContacts);
</script>

<AlertContainer {showAlert} {lastMesage} />
<div
	class="right-chat nav-wrap mt-2 right-scroll-bar"
	class:active-sidebar={show}
>
	<div class="middle-sidebar-right-content bg-white shadow-xss rounded-xxl">
		<!-- loader wrapper -->
		<div class="preloader-wrap p-3">
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
		</div>
		<!-- loader wrapper -->

		<div class="section full pe-3 ps-4 pt-4 position-relative feed-body">
			<h4 class="font-xsssss text-grey-500 text-uppercase fw-700 ls-3">
				CONTACTS
			</h4>
			<ul class="list-group list-group-flush">
				{#if Users && Users.length > 0}
					{#each Users as user}
						<li
							class="bg-transparent list-group-item no-icon pe-0 ps-0 pt-2 pb-2 border-0 d-flex align-items-center"
							on:click={toggleConstants}
						>
							<figure class="avatar float-left mb-0 me-2">
								<img src="/images/user-8.png" alt="display photo" class="w35" />
							</figure>
							<h3 class="fw-700 mb-0 mt-0">
								<a
									class="font-xssss text-grey-600 d-block text-dark model-popup-chat"
									href="/main/chat/{user?.ID}"
									>{user?.FirstName} {user?.LastName}</a
								>
							</h3>
							{#if user.Online == true}
								<span class="bg-success ms-auto btn-round-xss"></span>
							{:else}
								<span class="bg-warning ms-auto btn-round-xss"></span>
							{/if}
						</li>
					{/each}
				{/if}
				<!-- {#if offline && offline.length > 0}
                {#each offline as user}
                    <li class="bg-transparent list-group-item no-icon pe-0 ps-0 pt-2 pb-2 border-0 d-flex align-items-center" on:click={toggleConstants}>
                        <figure class="avatar float-left mb-0 me-2">
                            <img src="/images/user-8.png" alt="display photo" class="w35">
                        </figure>
                        <h3 class="fw-700 mb-0 mt-0">
                            <a class="font-xssss text-grey-600 d-block text-dark model-popup-chat" href="/main/chat/{user?.ID}">{user?.FirstName} {user?.LastName}</a>
                        </h3>

                        <span class="bg-warning ms-auto btn-round-xss"></span>
                    </li>
                {/each}
                {/if} -->
				<!-- <li class="bg-transparent list-group-item no-icon pe-0 ps-0 pt-2 pb-2 border-0 d-flex align-items-center">
                    <figure class="avatar float-left mb-0 me-2">
                        <img src="images/user-7.png" alt="display photo" class="w35">
                    </figure>
                    <h3 class="fw-700 mb-0 mt-0">
                        <a class="font-xssss text-grey-600 d-block text-dark model-popup-chat" href="/">Victor Exrixon</a>
                    </h3>
                    <span class="bg-success ms-auto btn-round-xss"></span>
                </li>
                <li class="bg-transparent list-group-item no-icon pe-0 ps-0 pt-2 pb-2 border-0 d-flex align-items-center">
                    <figure class="avatar float-left mb-0 me-2">
                        <img src="images/user-6.png" alt="display photo" class="w35">
                    </figure>
                    <h3 class="fw-700 mb-0 mt-0">
                        <a class="font-xssss text-grey-600 d-block text-dark model-popup-chat" href="/">Surfiya Zakir</a>
                    </h3>
                    <span class="bg-warning ms-auto btn-round-xss"></span>
                </li>
                <li class="bg-transparent list-group-item no-icon pe-0 ps-0 pt-2 pb-2 border-0 d-flex align-items-center">
                    <figure class="avatar float-left mb-0 me-2">
                        <img src="images/user-5.png" alt="display photo" class="w35">
                    </figure>
                    <h3 class="fw-700 mb-0 mt-0">
                        <a class="font-xssss text-grey-600 d-block text-dark model-popup-chat" href="/">Goria Coast</a>
                    </h3>
                    <span class="bg-success ms-auto btn-round-xss"></span>
                </li>
                <li class="bg-transparent list-group-item no-icon pe-0 ps-0 pt-2 pb-2 border-0 d-flex align-items-center">
                    <figure class="avatar float-left mb-0 me-2">
                        <img src="images/user-4.png" alt="display photo" class="w35">
                    </figure>
                    <h3 class="fw-700 mb-0 mt-0">
                        <a class="font-xssss text-grey-600 d-block text-dark model-popup-chat" href="/">Hurin Seary</a>
                    </h3>
                    <span class="badge mt-0 text-grey-500 badge-pill pe-0 font-xsssss">4:09 pm</span>
                </li>
                <li class="bg-transparent list-group-item no-icon pe-0 ps-0 pt-2 pb-2 border-0 d-flex align-items-center">
                    <figure class="avatar float-left mb-0 me-2">
                        <img src="images/user-3.png" alt="display photo" class="w35">
                    </figure>
                    <h3 class="fw-700 mb-0 mt-0">
                        <a class="font-xssss text-grey-600 d-block text-dark model-popup-chat" href="/">David Goria</a>
                    </h3>
                    <span class="badge mt-0 text-grey-500 badge-pill pe-0 font-xsssss">2 days</span>
                </li>
                <li class="bg-transparent list-group-item no-icon pe-0 ps-0 pt-2 pb-2 border-0 d-flex align-items-center">
                    <figure class="avatar float-left mb-0 me-2">
                        <img src="images/user-2.png" alt="display photo" class="w35">
                    </figure>
                    <h3 class="fw-700 mb-0 mt-0">
                        <a class="font-xssss text-grey-600 d-block text-dark model-popup-chat" href="/">Seary Victor</a>
                    </h3>
                    <span class="bg-success ms-auto btn-round-xss"></span>
                </li>
                <li class="bg-transparent list-group-item no-icon pe-0 ps-0 pt-2 pb-2 border-0 d-flex align-items-center">
                    <figure class="avatar float-left mb-0 me-2">
                        <img src="images/user-12.png" alt="display photo" class="w35">
                    </figure>
                    <h3 class="fw-700 mb-0 mt-0">
                        <a class="font-xssss text-grey-600 d-block text-dark model-popup-chat" href="/">Ana Seary</a>
                    </h3>
                    <span class="bg-success ms-auto btn-round-xss"></span>
                </li> -->
			</ul>
		</div>
	</div>
</div>

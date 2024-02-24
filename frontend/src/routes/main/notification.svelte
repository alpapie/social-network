<script>
	import axios from "axios";
	import { onMount } from "svelte";

	let NotifSocket;
	export let data;
	let showDot = false;

    const getNotifMessageSwitchType = (type,name) => {
		let message = "";
		switch (type) {
			case "follow":
				message = name+ " is following you.";
				break;
			case "invited-to-join-Group":
				message = name+ " accepted your follow request.";
				break;
			case "join-Group":
				message = name+ " declined your follow request.";
				break;
			default:
				message = "no message";
				break;
		}
		console.log("notification type:", type);
		return message;
	};

	const markAsRead = async (notifId, receiverId) => {
		try {
			let header = {
				cookie: document.cookie,
			};

    
			const config = {
				method: "get",
				withCredentials: true,
				header,
				mode: "no-cors",
				params: { user_id: receiverId, notif_id: notifId },
			};

			let httpResponse = await axios(
				`http://localhost:8080/server/notifAsRead`,
				config
			);
			console.log(httpResponse);
		} catch (error) {
			console.log(error);
		}
	};

	const respondToRequest = async (reply, notifId, receiverId) => {
		let httpResponse;
		let url = "";
		switch (reply) {
			case "accept":
				url = "acceptfollow";
				break;
			case "decline":
				url = "declinefollow";
				break;
			default:
				break;
		}

		try {
			let header = {
				cookie: document.cookie,
			};
			const config = {
				method: "get",
				withCredentials: true,
				header,
				mode: "no-cors",
				params: { user_id: receiverId, notif_id: notifId },
			};
			httpResponse = await axios(`http://localhost:8080/server/${url}`, config);
			console.log("reply response", httpResponse);
		} catch (error) {
			console.log(error);
		}
	};

	onMount(async () => {
		if (!NotifSocket) {
			NotifSocket = new WebSocket("ws://localhost:8080/server/initnotifsocket");
		}

		console.log("component data:", data);

		if (data.notifications.length > 0) {
			showDot = true;
		}

		console.log("initialise socket", NotifSocket);

		NotifSocket.onmessage = function (event) {
			let newEvent = JSON.parse(event.data);

			if (newEvent?.action == "notification") {
				data.notifications = [...data.notifications, newEvent.notification];
			}
			showDot = true;
		};

		NotifSocket.onclose = function (event) {
			NotifSocket = null;
		};

		NotifSocket.onerror = function (error) {
			NotifSocket = null;
		};
	});
</script>

<a
	href="#"
	class="p-2 text-center ms-auto menu-icon"
	id="dropdownMenu3"
	data-bs-toggle="dropdown"
	aria-expanded="false"
>
	{#if showDot}
		<span class="dot-count bg-warning"></span>
	{/if}
	<i class="feather-bell font-xl text-current"></i>
</a>
<div
	class="dropdown-menu dropdown-menu-end p-4 rounded-3 border-0 shadow-lg"
	aria-labelledby="dropdownMenu3"
>
	<h4 class="fw-700 font-xss mb-4">Notifications</h4>
	<!-- {#if !showDot}
		<div class="card bg-transparent-card w-100 border-0 ps-5 mb-3">
			<p>No new notifications</p>
		</div>
	{:else} -->
	{#each data.notifications as notif}
		<div class="card bg-transparent-card w-100 border-0 ps-5 mb-3">
			<img
				class="w40 position-absolute left-0"
				src="//ui-avatars.com/api/?name={notif.firstname +
					' ' +
					notif.lastname}&size=100&rounded=true&color=fff&background=random"
				width="50"
				alt={notif.firstName + " " + notif.lastName}
			/>
			<h5 class="font-xsss text-grey-900 mb-1 mt-0 fw-700 d-block">
				{notif.firstname + " " + notif.lastname}<span
					class="text-grey-400 font-xsssss fw-600 float-right mt-1"
				>
					3 min</span
				>
			</h5>
			<h6 class="text-grey-500 fw-500 font-xssss lh-4">
				{getNotifMessageSwitchType(notif.type)}
			</h6>
		</div>

		{#if notificationRequireConfirmation(notif.type)}
			<div class="card-body d-flex pt-0 ps-4 pe-4 pb-4 w50">
				<button
					on:click={() => {
						respondToRequest("accept", notif.id, notif.sender_id);
					}}
					class="p-2 w100 bg-success me-2 text-white text-center font-xssss fw-400 ls-1 rounded-xl"
					style="cursor: pointer;">Confirm</button
				>
				<button
					on:click={() => {
						respondToRequest("decline", notif.id, notif.sender_id);
					}}
					class="p-2 lh-20 text-white bg-danger w100 text-center font-xssss fw-400 ls-1 rounded-xl"
					style="cursor: pointer;">Decline</button
				>
			</div>
		{:else}
			<button
				on:click={() => {
					markAsRead(notif.id, notif.sender_id);
				}}
				class="p-2 lh-20 text-white bg-danger w100 text-center font-xssss fw-400 ls-1 rounded-xl"
				style="border: none;">Mark as read</button
			>
		{/if}
	{/each}
	<!-- {/if} -->
</div>

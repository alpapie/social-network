<script>
	import { onMount } from "svelte";
	import { page } from "$app/stores";
	import EmojiPicker from "svelte-emoji-picker";
	import { afterUpdate } from "svelte";
	import { LastMessage } from "../../stores";
	import { ContactsStore } from "../../stores";
	let Users;
	let usersOnline;
	ContactsStore.subscribe((val) => {
		Users = val;
	});

	import { WS } from "../../socket";
	export let data;
	let element;

	$: allmessage = data?.res?.result?.messages;
	afterUpdate(() => {
		if (allmessage) scrollToBottom(element);
	});
	$: if (allmessage && element) {
		scrollToBottom(element);
	}
	let socket;
	$: receiver_id = $page.params.id;
	let lastMesage = "";
	onMount(async () => {
		WS.subscribe((val) => (socket = val));
		socket.addEventListener("open", () => {
			console.log("Opened");
		});
		socket.onmessage = (e) => {
			var newMessage = JSON.parse(e.data);
			if (newMessage.action == "onlineUsers") {
				usersOnline = newMessage.users;
				for (let user of Users) {
					if (usersOnline.some((onlineUser) => onlineUser.ID === user.ID)) {
						user.Online = true;
					} else {
						user.Online = false;
					}
				}
				ContactsStore.set(Users);
			}
			if (
				(newMessage.sender_id == receiver_id ||
					newMessage.receiver_id == receiver_id) &&
				newMessage.groupId == 0
			) {
				if (allmessage) {
					allmessage = [...allmessage, newMessage];
				} else {
					allmessage = [newMessage];
				}
			}
			if (newMessage.sender_id == receiver_id) {
				LastMessage.set(newMessage);
			}
		};
	});
	const scrollToBottom = async (node) => {
		node.scroll({ top: node.scrollHeight, behavior: "smooth" });
	};
	let Message = "";

	let showEmojis = false;

	function SendMessage() {
		let mess = {
			receiver_id: Number(receiver_id),
			content: Message,
		};
		Message = "";
		const messageJSON = JSON.stringify(mess);
		socket.send(messageJSON);
	}
	function Send(e) {
		if (e.code == "Enter") {
			SendMessage();
		}
	}
</script>

<div class="main-content right-chat-active">
	<div class="middle-sidebar-bottom">
		<div
			class="middle-sidebar-left pe-0 ps-lg-3 ms-0 me-0"
			style="max-width: 100%;"
		>
			<div class="row">
				<div class="col-lg-12 position-relative">
					<div
						class="chat-wrapper pt-0 w-100 position-relative scroll-bar bg-white theme-dark-bg"
						bind:this={element}
						style="overflow:auto;"
					>
						<div class="chat-body p-3">
							<div class="messages-content pb-5">
								{#if allmessage?.length > 0}
									{#each allmessage as message}
										{#if message.receiver_id != receiver_id}
											<div class="message-item">
												<div class="message-user">
													<figure class="avatar">
														<img
															src="/images/user-9.png"
															alt={message.receiver_id}
														/>
													</figure>
													<div>
														<h5>{message.firstName} {message.lastName}</h5>
														<div class="time">01:35 PM</div>
													</div>
												</div>
												<div class="message-wrap">{message.content}</div>
											</div>
										{:else}
											<div class="message-item outgoing-message">
												<div class="message-user">
													<figure class="avatar">
														<img
															src="/images/user-1.png"
															alt="display p{message.receiver_id}"
														/>
													</figure>
													<div>
														<h5>{message.firstName} {message.lastName}</h5>
														<div class="time">
															01:35 PM<i class="ti-double-check text-info"></i>
														</div>
													</div>
												</div>
												<div class="message-wrap">{message.content}</div>
											</div>
										{/if}
									{/each}
								{/if}
								<div class="clearfix"></div>
							</div>
						</div>
					</div>
					<div
						class="chat-bottom dark-bg p-3 shadow-none theme-dark-bg"
						style="width: 98%;"
					>
						<div class="chat-form">
							<div style="display: {showEmojis ? 'block' : 'none'};">
								<EmojiPicker bind:value={Message} />
							</div>

							<button
								class="bg-current float-left"
								on:click={() => (showEmojis = !showEmojis)}>😀</button
							>

							<div class="form-group">
								<input
									type="text"
									placeholder="Start typing.."
									bind:value={Message}
									style="color: black;"
									on:keydown={Send}
								/>
							</div>
							<button class="bg-current" on:click={SendMessage}
								><i class="ti-arrow-right text-white"></i></button
							>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>

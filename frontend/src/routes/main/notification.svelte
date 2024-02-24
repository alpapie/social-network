<script>
    import { onMount } from "svelte";

    let NotifSocket
    export let data

    const getNotifMessageSwitchType = (name, type) => {
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
	const respondToRequest = (response, userId) => {
		switch (response) {
			case "accept":
				console.log(response, userId);
				break;
			case "decline":
				console.log(response, userId);
				break;
			default:
				break;
		}
	};

    onMount( async ()=>{
        if (!NotifSocket) {
            NotifSocket=  new WebSocket("ws://localhost:8080/server/initnotifsocket")
        }
        console.log("initialise socket",NotifSocket)

        NotifSocket.onmessage = function(event) {
            alert(`[message] Data received from server: ${event.data}`);
            console.log("new event", event);
			alert(`[message] Data received from server: ${event.data}`);
			let incomingNotification = JSON.parse(event.data);
			console.log("new notification message:", incomingNotification);
			data.notifications.push(incomingNotification);
        };

        NotifSocket.onclose = function(event) {
            NotifSocket=null
        };

        NotifSocket.onerror = function(error) {
            NotifSocket=null
        };
        
    })

</script>
<a href="#" class="p-2 text-center ms-auto menu-icon" id="dropdownMenu3" data-bs-toggle="dropdown" aria-expanded="false"><span class="dot-count bg-warning"></span><i class="feather-bell font-xl text-current"></i></a>
<div class="dropdown-menu dropdown-menu-end p-4 rounded-3 border-0 shadow-lg" aria-labelledby="dropdownMenu3">

    <h4 class="fw-700 font-xss mb-4">Notification</h4>
    {#each data.notifications as notif}
        <div class="card bg-transparent-card w-100 border-0 ps-5 mb-3">
            <img class="w40 position-absolute left-0" src="//ui-avatars.com/api/?name={notif.firstname+' '+ notif.lastname}&size=10&rounded=true&color=fff&background=random" width="50" alt="{notif.FirstName +" "+ notif.LastName}" />
            <h5 class="font-xsss text-grey-900 mb-1 mt-0 fw-700 d-block">{notif.firstname+' '+ notif.lastname}<span class="text-grey-400 font-xsssss fw-600 float-right mt-1"> 3 min</span></h5>
            <h6 class="text-grey-500 fw-500 font-xssss lh-4">{getNotifMessageSwitchType(notif.type)}</h6>
        </div>
        <div class="card-body d-flex pt-0 ps-4 pe-4 pb-4 w50">
            <a href="#" on:click={() => {
                respondToRequest("accept", notif.lastName);
            }} class="p-2 w100 bg-success me-2 text-white text-center font-xssss fw-400 ls-1 rounded-xl">Confirm</a>
            <a href="#" on:click={() => {
                respondToRequest("decline", notif.lastName);
            }}
            class="p-2 lh-20 text-white bg-danger w100 text-center font-xssss fw-400 ls-1 rounded-xl">Decline</a>
        </div>
    {/each}


</div>
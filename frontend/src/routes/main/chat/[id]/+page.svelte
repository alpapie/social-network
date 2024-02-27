<script>
    import {onMount} from "svelte"
    import { enhance } from "$app/forms";
    import { page } from "$app/stores"
    import EmojiPicker from 'svelte-emoji-picker';
    import { afterUpdate, tick } from 'svelte';

    import {WS} from "../../socket"
    export let data
    let element;
    
    $:console.log('data GUEYE receiver', data);
    $:allmessage = data?.res?.result?.messages
    afterUpdate(() => {
            console.log("afterUpdate");
            if(allmessage) scrollToBottom(element);
    });
    $: if(allmessage && element) {
		console.log("tick");
		scrollToBottom(element);
	}
    let socket
    $:receiver_id = $page.params.id
    $:console.log('RECEIVER ID', receiver_id);
    onMount(async () => { 
        // scrollToBottom(element);
        WS.subscribe((val)=> socket = val)
        socket.addEventListener("open", ()=> {
            console.log("Opened")
        })
        socket.onmessage = (e) => {
            var newMessage = JSON.parse(e.data);
            console.log("receive new message ", newMessage)
            if ((newMessage.sender_id == receiver_id || newMessage.receiver_id == receiver_id ) && newMessage.groupId == 0) {
                if (allmessage) {
                    allmessage = [...allmessage , newMessage]
                } else {
                    allmessage = [newMessage]
                }
                
            }
            // scrollToBottom(element);
            
        }
    })
    const scrollToBottom = async (node) => {
        node.scroll({ top: node.scrollHeight, behavior: 'smooth' });
    };
    let Message = ''
    let showEmojis = false;

    function SendMessage() {
        console.log("THE NEW MESSAGE IS" , Message)
        let mess = {
            // sender_id: 9,
            receiver_id : Number(receiver_id),
            content : Message,
        }
        Message = ''
        const messageJSON = JSON.stringify(mess);
        socket.send(messageJSON)
    }
    function Send(e) {
        if (e.code == "Enter" ) {
            SendMessage()
        }
    }
    // let MessagesContainer
</script>

<div class="main-content right-chat-active">
            
    <div class="middle-sidebar-bottom">
        <div class="middle-sidebar-left pe-0 ps-lg-3 ms-0 me-0" style="max-width: 100%;">
            <div class="row">
                   

                <div class="col-lg-12 position-relative">
                    <div class="chat-wrapper pt-0 w-100 position-relative scroll-bar bg-white theme-dark-bg" bind:this={element} style="overflow:auto;">
                        <div class="chat-body p-3 " >
                            <div class="messages-content pb-5" >
                                {#if allmessage?.length > 0}
                                    {#each allmessage as message }
                                    {#if message.receiver_id != receiver_id}
                                        <div class="message-item" >
                                            <div class="message-user">
                                                <figure class="avatar">
                                                    <img src="/images/user-9.png" alt="display photo">
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
                                                    <img src="/images/user-1.png" alt="display photo">
                                                </figure>
                                                <div>
                                                    <h5>{message.firstName} {message.lastName}</h5>
                                                    <div class="time">01:35 PM<i class="ti-double-check text-info"></i></div>
                                                </div>
                                            </div>
                                            <div class="message-wrap">{message.content}</div>
                                        </div>
                                    {/if} 
                                   
                                    {/each}
                                {/if}
                                    

                                <!-- <div class="message-item">
                                    <div class="message-user">
                                        <figure class="avatar">
                                            <img src="/images/user-9.png" alt="display photo">
                                        </figure>
                                        <div>
                                            <h5>Byrom Guittet</h5>
                                            <div class="time">01:35 PM</div>
                                        </div>
                                    </div>
                                    <div class="message-wrap">I've found some cool photos for our travel app.</div>
                                </div>

                                <div class="message-item outgoing-message">
                                    <div class="message-user">
                                        <figure class="avatar">
                                            <img src="/images/user-1.png" alt="display photo">
                                        </figure>
                                        <div>
                                            <h5>Byrom Guittet</h5>
                                            <div class="time">01:35 PM<i class="ti-double-check text-info"></i></div>
                                        </div>
                                    </div>
                                    <div class="message-wrap">Hey mate! How are things going ?</div>
                                </div>

                                <div class="message-item">
                                    <div class="message-user">
                                        <figure class="avatar">
                                            <img src="/images/user-9.png" alt="display photo">
                                        </figure>
                                        <div>
                                            <h5>Byrom Guittet</h5>
                                            <div class="time">01:35 PM</div>
                                        </div>
                                    </div>
                                    <figure>
                                        <img src="/images/bb-9.jpg" class="w-25 img-fluid rounded-3" alt="display photo">
                                    </figure>
                                    
                                
                                </div>

                                <div class="message-item outgoing-message">
                                    <div class="message-user">
                                        <figure class="avatar">
                                            <img src="/images/user-1.png" alt="display photo">
                                        </figure>
                                        <div>
                                            <h5>Byrom Guittet</h5>
                                            <div class="time">01:35 PM<i class="ti-double-check text-info"></i></div>
                                        </div>
                                    </div>
                                    <div class="message-wrap" style="margin-bottom: 90px;">Hey mate! How are things going ?</div>

                                </div> -->
                                <div class="clearfix"></div>


                            </div>
                        </div>
                    </div>
                    <div class="chat-bottom dark-bg p-3 shadow-none theme-dark-bg" style="width: 98%;">
                        <div class="chat-form">
                            <div style="display: {showEmojis ? 'block' : 'none'};" >
                                <EmojiPicker bind:value={Message} />
                            </div>

                            <button class="bg-current float-left" on:click={() => showEmojis = !showEmojis}>😀</button>

                            <div class="form-group">
                                <input type="text" placeholder="Start typing.."  bind:value={Message} style="color: black;" on:keydown={Send}>
                            </div>          
                            <button class="bg-current" on:click={SendMessage}><i class="ti-arrow-right text-white" ></i></button>
                        </div>
                    </div> 
                </div>

            </div>
        </div>
         
    </div>            
</div>
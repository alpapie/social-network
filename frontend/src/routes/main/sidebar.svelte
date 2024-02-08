<!-- navigation left -->
<script>
  import { createGroup } from "$lib/groups/createGroup";
    import { Modal, Content, Trigger } from "sv-popup"
    import { writable } from 'svelte/store';
    let groupTitle = '';
    let groupDescription = '';
    $: closed =false
    async function handleSubmit() {
        console.log('Group Title:', groupTitle);
        console.log('Group Description:', groupDescription);
        const groupData = {
        title: groupTitle,
        description: groupDescription
      };      
      try {
      const response = await createGroup(groupData);
      console.log('Group created successfully:', response);
      closed = true;
    } catch (error) {
      console.error('Failed to create group:', error);
    }
        // closed=true
    }
</script>

<nav class="navigation scroll-bar">
    <div class="container ps-0 pe-0">
        <div class="nav-content">
            <div class="nav-wrap bg-white bg-transparent-card rounded-xxl shadow-xss pt-3 pb-1 mb-2 mt-2">
                <div class="nav-caption fw-600 font-xssss text-grey-500"><span>New </span>Feeds</div>
                <ul class="mb-1 top-content">
                    <li class="logo d-none d-xl-block d-lg-block"></li>
                    <Modal button={false} close={closed}>
                        <Content>
                          <form on:submit|preventDefault={handleSubmit}>
                            <input id="groupTitle" bind:value={groupTitle} required placeholder="Title"/>
                            <textarea id="groupDescription" bind:value={groupDescription} required placeholder="Description"></textarea>
                            <button type="submit" >Create</button>
                          </form>
                        </Content>
                        <Trigger>
                        <li><a href="#" class="nav-content-bttn open-font" ><i class="feather-tv btn-round-md bg-blue-gradiant me-3"></i><span>New Group</span></a></li>                            
                        </Trigger>
                      </Modal>
                    <li><a href="default-storie.html" class="nav-content-bttn open-font" ><i class="feather-globe btn-round-md bg-gold-gradiant me-3"></i><span>Explore Stories</span></a></li>
                    <li><a href="/groups" class="nav-content-bttn open-font" ><i class="feather-zap btn-round-md bg-mini-gradiant me-3"></i><span>Popular Groups</span></a></li>
                    <li><a href="user-page.html" class="nav-content-bttn open-font"><i class="feather-user btn-round-md bg-primary-gradiant me-3"></i><span>Author Profile </span></a></li>                        
                </ul>
            </div>

            <div class="nav-wrap bg-white bg-transparent-card rounded-xxl shadow-xss pt-3 pb-1 mb-2">
                <div class="nav-caption fw-600 font-xssss text-grey-500"><span>More </span>Pages</div>
                <ul class="mb-3">
                    <li><a href="default-email-box.html" class="nav-content-bttn open-font"><i class="font-xl text-current feather-inbox me-3"></i><span>Email Box</span><span class="circle-count bg-warning mt-1">584</span></a></li>
                    <li><a href="default-hotel.html" class="nav-content-bttn open-font"><i class="font-xl text-current feather-home me-3"></i><span>Near Hotel</span></a></li>
                    <li><a href="default-event.html" class="nav-content-bttn open-font"><i class="font-xl text-current feather-map-pin me-3"></i><span>Latest Event</span></a></li>
                    <li><a href="default-live-stream.html" class="nav-content-bttn open-font"><i class="font-xl text-current feather-youtube me-3"></i><span>Live Stream</span></a></li>                        
                </ul>
            </div>
            <div class="nav-wrap bg-white bg-transparent-card rounded-xxl shadow-xss pt-3 pb-1">
                <div class="nav-caption fw-600 font-xssss text-grey-500"><span></span> Account</div>
                <ul class="mb-1">
                    <li class="logo d-none d-xl-block d-lg-block"></li>
                    <li><a href="default-settings.html" class="nav-content-bttn open-font h-auto pt-2 pb-2"><i class="font-sm feather-settings me-3 text-grey-500"></i><span>Settings</span></a></li>
                    <li><a href="default-analytics.html" class="nav-content-bttn open-font h-auto pt-2 pb-2"><i class="font-sm feather-pie-chart me-3 text-grey-500"></i><span>Analytics</span></a></li>
                    <li><a href="default-message.html" class="nav-content-bttn open-font h-auto pt-2 pb-2"><i class="font-sm feather-message-square me-3 text-grey-500"></i><span>Chat</span><span class="circle-count bg-warning mt-0">23</span></a></li>
                </ul>
            </div>
        </div>
    </div>
</nav>

<style>
    form {
      display: flex;
      flex-direction: column;
      gap:  1rem;
      max-width:  400px;
      margin: auto;
      padding:  1rem;
      border:  1px solid #ccc;
      border-radius:  5px;
      background-color: #f9f9f9;
          display: flex;
          justify-content: center;
          align-items: center;
    }
  
    label {
      font-weight: bold;
      margin-bottom:  0.5rem;
    }
  
    input,
    textarea {
      padding:  0.5rem;
      border:  1px solid #ccc;
      border-radius:  3px;
      width:  100%;
    }
  
    textarea {
      resize: vertical;
      min-height:  100px;
    }
  
    button {
      padding:  0.5rem  1rem;
      background-color: #007bff;
      color: white;
      border: none;
      border-radius:  3px;
      cursor: pointer;
      transition: background-color  0.3s ease;
    }
  
    button:hover {
      background-color: #0056b3;
    }
  </style>
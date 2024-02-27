<!-- navigation left -->
<script>
    import { enhance } from "$app/forms";
    import { Modal, Content, Trigger } from "sv-popup"

    $: closed = false
    export let data
    let user=JSON.parse(data)
</script>

<nav class="navigation scroll-bar">
    <div class="container ps-0 pe-0">
        <div class="nav-content">
          <div class="card" style="border-radius: 15px;margin-top:20px">
            <div class="card-body text-center">
              <div class="mt-3 mb-4">
                {#if !user?.Avatar}
                <img src="{user?.Avatar}"
                  class="rounded-circle img-fluid" style="width: 100px;" />
                {:else}
                <img src="//ui-avatars.com/api/?name={user.FirstName +" "+ user.LastName}&size=50&rounded=true&color=fff&background=random"
                  class="rounded-circle img-fluid" style="width: 100px;" />
                {/if}
              </div>
              <h4 class="mb-2"> {user?.FirstName+" "+user?.LastName}</h4>
              <p class="text-muted mb-4">@{user?.NickName} <span class="mx-2"></span> <a
                  href="#!">{user?.Email}</a></p>
            </div>
          </div>
            <div class="nav-wrap bg-white bg-transparent-card rounded-xxl shadow-xss pt-3 pb-1 mb-2 mt-2">
              <div class="nav-caption fw-600 font-xssss text-grey-500"><span>New </span>Feeds</div>
              
                <ul class="mb-1 top-content">
                    <li class="logo d-none d-xl-block d-lg-block"></li>
                    <Modal button={false} close={closed}>
                        <Content>
                          <form method="post"  action="/main?/createGroup" use:enhance on:submit={()=>{closed=true}}>
                            <input id="groupTitle" name="title" required placeholder="Title"/>
                            <textarea id="groupDescription" name="description" required placeholder="Description"></textarea>
                            <button type="submit" on:click={()=>{closed=false}}>Create</button>
                          </form>
                        </Content>
                        <Trigger>
                        <li><a href="#" class="nav-content-bttn open-font" ><i class="feather-tv btn-round-md bg-blue-gradiant me-3"></i><span>New Group</span></a></li>                            
                        </Trigger>
                      </Modal>
                    <li><a href="default-storie.html" class="nav-content-bttn open-font" ><i class="feather-globe btn-round-md bg-gold-gradiant me-3"></i><span>Explore Stories</span></a></li>
                    <li><a href="/main/groups" class="nav-content-bttn open-font" ><i class="feather-zap btn-round-md bg-mini-gradiant me-3"></i><span>Popular Groups</span></a></li>
                    <li><a href="/main/profile" class="nav-content-bttn open-font"><i class="feather-user btn-round-md bg-primary-gradiant me-3"></i><span>Author Profile </span></a></li>                        
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

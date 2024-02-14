<script>
    import { Modal, Content, Trigger } from "sv-popup"
    import { makeRequest } from "$lib/api";
    
    export let data;
    console.log(data);
    let groups = data.res.result;
    let filter = 'joined';
    

    function updateFilter(newFilter) {
        filter = newFilter;
    }
    $: closed =false

    function validateForm(groupIdInput, userIdInput, usersSelect) {

        if (isNaN(groupIdInput) || isNaN(userIdInput)) {
            return false;
        }

        if (usersSelect.length ===  0) {
            return false;
        }

        return true;
    }

    let selectedUsers = [];
    let formerror = '';
    let copiedGroups = {...groups};
    async function handleSubmit() {
        const groupData = {
            groupId:  document.getElementById('groupId').value,
            userId:  document.getElementById('userId').value,
            userIds: selectedUsers
        };

        try {
            if (validateForm(groupData.groupId, groupData.userId, groupData.userIds)) {
                const response = await makeRequest("invitegroup","POST",JSON.stringify(groupData))

                if (response) {
                    throw new Error(`HTTP error! status:`);
                }
                closed = true;
                formerror = '';
            }else{
                formerror = 'ERROR ON FORM'
            }
        } catch (error) {
            console.error('ERROR : ', error);
        }
    }

    // let copiedGroups = {...groups};
    async function handleFollowClick(groupId, userId) {
        console.log('CLICKED');
        const requestData = {
            groupId: groupId,
            userId: userId
        };

        try {
            const response = await makeRequest('followgroup', 'POST', JSON.stringify(requestData));

            if (response) {
                throw new Error(`HTTP error! status:`);
            }
            console.log('REQUEST SENDED');
            copiedGroups.Notjoined = copiedGroups.Notjoined.map(group => {
                if (group.id === groupId) {
                    return {...group, notif_type: true};
                }
                return group;
            });

        } catch (error) {
            console.error('ERROR : ', error);
        }
    }

    let searchValue = '';
    function filterUsers(searchText, indice) {
        return groups.joined[indice].suggests.filter(user => {
            return user.first_name.toLowerCase().includes(searchText.toLowerCase()) ||
                user.last_name.toLowerCase().includes(searchText.toLowerCase());
        });
    }

    
    
    
</script>

<div class="main-content right-chat-active">
    
            
    <div class="middle-sidebar-bottom">
        
        <div class="middle-sidebar-left pe-0">
            <div class="row">
                <div class="col-xl-12">
                    <div class="card shadow-xss w-100 d-block d-flex border-0 p-4 mb-3">
                        <div class="card-body d-flex justify-content-between align-items-center p-0">
                            <h2 class="fw-700 mb-0 mt-0 font-md text-grey-900">Group</h2>
                            <div class="d-flex">
                                <button class="btn btn-custom ml-2 {filter === 'joined' ? 'active' : ''}" on:click={() => updateFilter('joined')}>Your Groups</button>
                                <button class="btn btn-custom ml-2 {filter === 'not' ? 'active' : ''}" on:click={() => updateFilter('not')}>Suggestions</button>
                            </div>
                        </div>
                    </div>

                    
                    <div class="row ps-2 pe-1" style="height: fit-content;">
                        {#if filter === 'joined'}
                        {#if groups.joined && groups.joined.length >  0}
                        {#each groups.joined as group, index}
                        
                        <div class="col-md-6 col-sm-6 pe-2 ps-2">
                            <div class="card d-block border-0 shadow-xss rounded-3 overflow-hidden mb-3">
                                <div class="card-body position-relative h100 bg-image-cover bg-image-center" style="background-image: url(images/bb-16.png);"></div>
                                <div class="card-body d-block w-100 pl-10 pe-4 pb-4 pt-0 text-left position-relative">
                                    <figure class="avatar position-absolute w75 z-index-1" style="top:-40px; left: 15px;"><img src="images/user-12.png" alt="image" class="float-right p-1 bg-white rounded-circle w-100"></figure>
                                    <div class="clearfix"></div>
                                    <a href="/main/groups/{group.id}">
                                    <h4 class="fw-700 font-xsss mt-3 mb-1">{group.title}</h4>
                                    <p class="fw-500 font-xsssss text-grey-500 mt-0 mb-3">{group.description}</p>
                                    </a>
                                    <span class="position-absolute right-15 top-0 d-flex align-items-center">
                                        <a href="#" class="d-lg-block d-none"><i class="feather-video btn-round-md font-md bg-primary-gradiant text-white"></i></a>
                                        <Modal button={false} close={closed}>
                                            <Content>       
                                                <form on:submit|preventDefault={handleSubmit}>
                                                    <div>Select users to invite into GROUP : {group.title}</div>
                                                    <input type="text" placeholder="Recherche" bind:value={searchValue} on:input={e => searchValue = e.target.value} />
                                                    <p class="danger fw-100 text-danger">
                                                        {#if formerror}
                                                           {formerror}
                                                        {/if}
                                                    </p>
                                                    <input type="hidden" id="groupId" name="groupId" value={group.id}>
                                                    <input type="hidden" id="userId" name="userId" value={groups.userid}>
                                                    <select id="usersSelect" name="usersSelect" multiple bind:value={selectedUsers}>
                                                        {#if groups.joined[0].suggests && groups.joined[0].suggests.length >  0}
                                                        {#each filterUsers(searchValue, index) as user}
                                                        {#if user.is_requested}
                                                            <option data-email="already requested" value="{user.id}" disabled>
                                                                {user.first_name} {user.last_name}
                                                            </option>
                                                        {:else}
                                                            <option data-email="{user.email}" value="{user.id}">
                                                                {user.first_name} {user.last_name}
                                                            </option>
                                                        {/if}
                                                        {/each}
                                                        {/if}
                                                    </select>
                                                    <button type="submit">Invite user(s)</button>
                                                </form>
                                            </Content>
                                            <Trigger>
                                                <a href="#" class="text-center p-2 lh-24 w100 ms-1 ls-3 d-inline-block rounded-xl bg-current font-xsssss fw-700 ls-lg text-white">INVITE</a>
                                            </Trigger>
                                        </Modal>
                                        
                                    </span>
                                </div>
                            </div>
                        </div>
                       
                        {/each}
                        {/if}
                        {:else if filter === 'not'}
                        {#if groups.Notjoined && groups.Notjoined.length >  0}
                        {#each copiedGroups.Notjoined as group}
                        <div class="col-md-6 col-sm-6 pe-2 ps-2">
                            <div class="card d-block border-0 shadow-xss rounded-3 overflow-hidden mb-3">
                                <div class="card-body position-relative h100 bg-image-cover bg-image-center" style="background-image: url(images/bb-16.png);"></div>
                                <div class="card-body d-block w-100 pl-10 pe-4 pb-4 pt-0 text-left position-relative">
                                    <figure class="avatar position-absolute w75 z-index-1" style="top:-40px; left: 15px;"><img src="images/user-12.png" alt="image" class="float-right p-1 bg-white rounded-circle w-100"></figure>
                                    <div class="clearfix"></div>
                                    <h4 class="fw-700 font-xsss mt-3 mb-1">{group.title}</h4>
                                    <p class="fw-500 font-xsssss text-grey-500 mt-0 mb-3">{group.description}</p>
                                    <span class="position-absolute right-15 top-0 d-flex align-items-center">
                                        <a href="#" class="d-lg-block d-none"><i class="feather-video btn-round-md font-md bg-primary-gradiant text-white"></i></a>
                                            {#if group.notif_type}
                                                <a id="followButton" class="text-center p-2 lh-24 w100 ms-1 ls-3 d-inline-block rounded-xl bg-grey font-xsssss fw-700 ls-lg text-grey-700">REQUESTED</a>
                                            {:else}
                                                <a id="followButton" style="cursor: pointer;" class="text-center p-2 lh-24 w100 ms-1 ls-3 d-inline-block rounded-xl bg-current font-xsssss fw-700 ls-lg text-white" on:click={() => handleFollowClick(group.id, groups.userid)}>FOLLOW</a>
                                            {/if}
                                    </span>
                                </div>
                            </div>
                        </div>
                        {/each}
                        {/if}
                        {/if}
                        
                    
                          
                    </div>
                </div>               
            </div>
        </div>
         
    </div>            
</div>

<style>
    #usersSelect {
        width: 400px;
        height: 200px;
    }

    #usersSelect option {
        padding:  0.5rem; 
    }
    

    .btn-custom {
        background-color: #fff; 
        color: #272727; 
        border: none; 
        transition: none; 
    }

    .btn-custom.active {
        color: blue;
    }
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

        #usersSelect option[data-email]:after {
        content: attr(data-email);
        display: block;
        font-size:  0.8em;
        color: #888;
    }
      

</style>
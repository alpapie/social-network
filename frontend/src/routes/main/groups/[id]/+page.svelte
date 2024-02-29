<script>
    import { Modal, Content, Trigger } from "sv-popup"
    import axios from "axios";
    import { enhance } from "$app/forms";
    import Posts from "../../post.svelte";
    import CreatePost from "./createpost.svelte"
    import Comment from "../../comments.svelte"
    export let data;
    console.log('data received GET', data);
    let grpInfo = data?.res?.result?.groupDetail?.groupdata;
    let evts = data?.res?.result?.groupDetail?.events
    console.log('event', evts);
    let nbrFollower = data?.res?.result?.groupDetail?.nbrfollowers
    export let form;
    $: closed =false
    $: console.log('data received', form);

    $: if (form?.error === 'no') {
        closed = true;
    }

    function getMonthName(monthNumber) {
        const months = ['JAN', 'FEB', 'MAR', 'APR', 'MAY', 'JUN', 'JUL', 'AUG', 'SEP', 'OCT', 'NOV', 'DEC'];
        return months[monthNumber -  1] || 'NaN';
    }
    function getEventColorClass(index) {
        const colors = ['bg-success', 'bg-warning', 'bg-primary'];
        return colors[index % colors.length];
    }
    let CommSection = { display: "none", data: {} };

    let responses = {}; 

    async function handleResponse(eventId, response) {
        let url = `http://localhost:8080/server/getoption?eventid=${eventId}&response=${response}`;
        try {
            let header={
                cookie:document.cookie
            }
            const config = { method:"get",withCredentials: true , header,mode: 'no-cors' };
            let responsess = await axios(url, config)
            console.log("SUCCES",  responsess.data);
            responses[eventId] = response;

            // return response.data
        } catch (err) {
            console.log("ERORR",  err);
        }
    }

       
</script>



<div class="main-content right-chat-active">
            
    <div class="middle-sidebar-bottom">
        <div class="middle-sidebar-left">
            <div class="row">
                <div class="col-xl-4 col-xxl-3 col-lg-4 pe-0">
                    <div class="card w-100 shadow-xss rounded-xxl overflow-hidden border-0 mb-3 mt-3 pb-3">
                        <div class="card-body position-relative h150 bg-image-cover bg-image-center" style="background-image: url(/images/bb-9.jpg);"></div>
                        <div class="card-body d-block pt-4 text-center">
                            <figure class="avatar mt--6 position-relative w75 z-index-1 w100 z-index-1 ms-auto me-auto"><img src="/images/pt-1.jpg" alt="image" class="p-1 bg-white rounded-xl w-100"></figure>
                            <h4 class="font-xs ls-1 fw-700 text-grey-900">{grpInfo?.title}<span class="d-block font-xssss fw-500 mt-1 lh-3 text-grey-500">@surfiyazakir22</span></h4>
                        </div>
                        <div class="card-body d-flex align-items-center ps-4 pe-4 pt-0">
                            <h4 class="font-xsssss text-center text-grey-500 fw-600 ms-2 me-2"><b class="text-grey-900 mb-1 font-xss fw-700 d-inline-block ls-3 text-dark">456 </b> Posts</h4>
                            <h4 class="font-xsssss text-center text-grey-500 fw-600 ms-2 me-2"><b class="text-grey-900 mb-1 font-xss fw-700 d-inline-block ls-3 text-dark">{nbrFollower} </b> Followers</h4>
                            <h4 class="font-xsssss text-center text-grey-500 fw-600 ms-2 me-2"><b class="text-grey-900 mb-1 font-xss fw-700 d-inline-block ls-3 text-dark">32k </b> Follow</h4>
                        </div>
                        <div class="card-body d-flex align-items-center justify-content-center ps-4 pe-4 pt-0">
                            <Modal button={false} close={closed}>
                                <Content>       
                                    <form action="?/createEvent" method="POST" use:enhance>
                                        <div>Create Your Event 🎰 </div>
                                        <input type="text" value={form?.title??''} id="event-title" name="title" placeholder="Title" required>
                                        <textarea id="event-description" value={form?.description??''} name="description" placeholder="description" required></textarea>
                                        <input type="date" value={form?.date??''} id="event-date" name="date" required>
                                        <input type="time" value={form?.time??'08:00'} id="event-time" name="time" required>
                                        {#if form && form?.error != 'no'}
                                            <div class="alert alert-danger">
                                                {form.error}
                                            </div>
                                        {/if}
                                        <button type="submit">Create</button>
                                    </form>
                                </Content>
                                <Trigger>
                                    <a href="#" class="bg-success p-3 z-index-1 rounded-3 text-white font-xsssss text-uppercase fw-700 ls-3">New Event</a>
                                </Trigger>
                            </Modal>
                            
                            <a href={`/main/groups/${grpInfo?.id}/messages`} class="bg-greylight btn-round-lg ms-2 rounded-3 text-grey-700"><i class="feather-mail font-md"></i></a>
                            <a href="#" class="bg-greylight theme-white-bg btn-round-lg ms-2 rounded-3 text-grey-700"><i class="ti-more font-md"></i></a>
                        </div>
                    </div>
                    <div class="card w-100 shadow-xss rounded-xxl border-0 mb-3">
                        <div class="card-body d-block p-4">
                            <h4 class="fw-700 mb-3 font-xsss text-grey-900">About</h4>
                            <p class="fw-500 text-grey-500 lh-24 font-xssss mb-0">{grpInfo?.description}</p>
                        </div>
                        <div class="card-body border-top-xs d-flex">
                            <i class="feather-lock text-grey-500 me-3 font-lg"></i>
                            <h4 class="fw-700 text-grey-900 font-xssss mt-0">Private <span class="d-block font-xssss fw-500 mt-1 lh-3 text-grey-500">What's up, how are you?</span></h4>
                        </div>

                        <div class="card-body d-flex pt-0">
                            <i class="feather-eye text-grey-500 me-3 font-lg"></i>
                            <h4 class="fw-700 text-grey-900 font-xssss mt-0">Visble <span class="d-block font-xssss fw-500 mt-1 lh-3 text-grey-500">Anyone can find you</span></h4>
                        </div>
                        <div class="card-body d-flex pt-0">
                            <i class="feather-map-pin text-grey-500 me-3 font-lg"></i>
                            <h4 class="fw-700 text-grey-900 font-xssss mt-1">Flodia, Austia </h4>
                        </div>
                        <div class="card-body d-flex pt-0">
                            <i class="feather-users text-grey-500 me-3 font-lg"></i>
                            <h4 class="fw-700 text-grey-900 font-xssss mt-1">Genarel Group</h4>
                        </div>
                    </div>
                    <div class="card w-100 shadow-xss rounded-xxl border-0 mb-3">
                        <div class="card-body d-flex align-items-center  p-4">
                            <h4 class="fw-700 mb-0 font-xssss text-grey-900">Photos</h4>
                            <a href="#" class="fw-600 ms-auto font-xssss text-primary">See all</a>
                        </div>
                        <div class="card-body d-block pt-0 pb-2">
                            <div class="row">
                                <div class="col-6 mb-2 pe-1"><a href="images/e-2.jpg" data-lightbox="roadtrip"><img src="/images/e-2.jpg" alt="image" class="img-fluid rounded-3 w-100"></a></div>
                                <div class="col-6 mb-2 ps-1"><a href="images/e-3.jpg" data-lightbox="roadtrip"><img src="/images/e-3.jpg" alt="image" class="img-fluid rounded-3 w-100"></a></div>
                                <div class="col-6 mb-2 pe-1"><a href="images/e-4.jpg" data-lightbox="roadtrip"><img src="/images/e-4.jpg" alt="image" class="img-fluid rounded-3 w-100"></a></div>
                                <div class="col-6 mb-2 ps-1"><a href="images/e-5.jpg" data-lightbox="roadtrip"><img src="/images/e-5.jpg" alt="image" class="img-fluid rounded-3 w-100"></a></div>
                                <div class="col-6 mb-2 pe-1"><a href="images/e-2.jpg" data-lightbox="roadtrip"><img src="/images/e-2.jpg" alt="image" class="img-fluid rounded-3 w-100"></a></div>
                                <div class="col-6 mb-2 ps-1"><a href="images/e-1.jpg" data-lightbox="roadtrip"><img src="/images/e-1.jpg" alt="image" class="img-fluid rounded-3 w-100"></a></div>
                            </div>
                        </div>
                        <div class="card-body d-block w-100 pt-0">
                            <a href="#" class="p-2 lh-28 w-100 d-block bg-grey text-grey-800 text-center font-xssss fw-700 rounded-xl"><i class="feather-external-link font-xss me-2"></i> More</a>
                        </div>
                    </div>

                    <div class="card w-100 shadow-xss rounded-xxl border-0 mb-3" style="overflow: auto;">
                        
                        <div class="card-body d-flex align-items-center  p-4">
                            <h4 class="fw-700 mb-0 font-xssss text-grey-900">Event</h4>
                            <a href="#" class="fw-600 ms-auto font-xssss text-primary">See all</a>
                        </div>
                        <!-- <div>{data.res.result.groupDetail.events}</div> -->
                        {#if data?.res?.result?.groupDetail?.events && data?.res?.result?.groupDetail?.events.length >  0}
                            {#each data?.res?.result?.groupDetail?.events as event, index}
                            <div class="event-card">
                                <div class="card-body d-flex pt-0 ps-4 pe-4 pb-3 overflow-hidden">
                                    <div class="{getEventColorClass(index)} me-2 p-3 rounded-xxl"><h4 class="fw-700 font-lg ls-3 lh-1 text-white mb-0">
                                        <span class="ls-1 d-block font-xsss text-white fw-600">
                                            {getMonthName(new Date(event.date).getMonth() +  1)}
                                        </span>
                                        {new Date(event.date).getDate().toString().padStart(2, '0')}</h4>
                                        <span class="ls-1 d-block font-xsss text-white fw-600">
                                            {new Date(event.date).getFullYear()}
                                        </span></div>
                                    <h4 class="fw-700 text-grey-900 font-xssss mt-2">{event.title}<span class="d-block font-xsssss fw-500 mt-1 lh-4 text-grey-500">{event.description}</span> </h4>
                                    
                                </div>
                                
                                <div class="fw-700 text-grey-900 font-xssss mt-2 question-container">
                                    <p>Voulez-vous participer à cet événement ?</p>
                                    <div class="d-flex justify-content-center" style="margin-bottom: 5%;">

                                        {#if responses[event.id] === 'Oui' || event.isgoing == true}
                                             <button class="btn bg-success mr-12 btn-full-width fw-700 text-grey-900 font-xssss mt-2" >You're IN 📌✔</button>
                                        {:else if responses[event.id] === 'Non' || event.isgoing == false}
                                        <button class="btn bg-warning mr-12 btn-full-width fw-700 text-grey-900 font-xssss mt-2" >Ya kham Gayin ❌</button>
                                        {:else}
                                        <button class="btn bg-success  mr-12  fw-700 text-grey-900 font-xssss mt-2" on:click={() => handleResponse(event.id, 'Oui')}>Oui</button>
                                        <button class="btn bg-grey  fw-700 text-grey-900 font-xssss mt-2" on:click={() => handleResponse(event.id, 'Non')}>Non</button>
                                        {/if}
                                    </div>
                                </div>
                            </div>
                            {/each}
                        {/if}
                         
                    </div>
                </div>
                <div class="col-xl-8 col-xxl-9 col-lg-8">
                    <CreatePost/>
                    <Posts data={data?.res?.result?.groupDetail?.postsdata} bind:CommSection/>
                </div>                             
            </div>
        </div>
         
    </div>            
</div>
<Comment bind:CommSection />


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

    .question-container {
        text-align: center;
        margin-top:  20px; 
    }
    .question-container p {
        margin-bottom:  10px; 
    }
    .question-container .btn {
        margin-right:  5px; 
    }
    .btn-full-width {
        width:  100%;
    }

    .event-card {
    background-color: white;
    border-radius:  10px;
    box-shadow:  0  4px  6px rgba(0,  0,  0,  0.1);
    margin-bottom:  20px;
    padding:  20px;
    transition: transform  0.3s ease;
}


.event-card h4 {
    margin-bottom:  10px;
}

.event-card p {
    margin-bottom:  10px;
}



</style>
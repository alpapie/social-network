<script>
  import Createpost from "../createpost.svelte";
  import Eventsugest from "../eventsugest.svelte";
  import Posts from "../posts.svelte";
  import axios from "axios";
  import UserList from "./userList.svelte";

  export let data;

  export let toprivate;
  let success = true;
  let erroralert = true;
  let privatestatus = data?.user?.IsPublic == 0;
  async function changeCountStatus(e) {
    toprivate = !toprivate;
    setTimeout(() => {
      toprivate = !toprivate;
    }, 60000);

    let header = {
      cookie: document.cookie,
    };
    const config = {
      method: "get",
      withCredentials: true,
      header,
      mode: "no-cors",
      params: { ispublic: privatestatus ? 1 : 0 },
    };
    let response = await axios(
      "http://localhost:8080/server/changestatus",
      config
    );
    if (!response?.data?.success) {
      erroralert = false;
      setTimeout(() => {
        erroralert = true;
      }, 2000);
      return;
    }
    success = false;
    setTimeout(() => {
      success = true;
    }, 2000);
  }
</script>

<!-- main content -->
<div class="main-content right-chat-active">
  <div class="middle-sidebar-bottom">
    <div class="middle-sidebar-left">
      <div
        id="success-alert"
        class="alert alert-success alert-dismissible {success
          ? 'd-none'
          : ''} fade show"
        role="alert"
      >
        status profil changed success!
      </div>
      <!-- Alert Error -->
      <div
        id="error-alert"
        class="alert alert-danger alert-dismissible {erroralert
          ? 'd-none'
          : ''} fade show"
        role="alert"
      >
        Error for change prfil status!
      </div>
      <div class="row">
        <div class="col-lg-12">
          <div class="card w-100 border-0 p-0 bg-white shadow-xss rounded-xxl">
            <div class="card-body h250 p-0 rounded-xxl overflow-hidden m-3">
              <img src="/images/u-bg.jpg" alt="image" />
            </div>
            <div class="card-body p-0 position-relative">
              <figure
                class="avatar position-absolute w100 z-index-1"
                style="top:-40px; left: 30px;"
              >
                {#if data?.user?.Avatar !== "''"}
                  <img
                    src={data?.user?.Avatar}
                    class="float-right p-1 bg-white rounded-circle w-100"
                    alt={data?.user?.FirstName}
                  />
                {:else}
                  <img
                    src="//ui-avatars.com/api/?name={data?.user?.FirstName +
                      ' ' +
                      data?.user
                        ?.LastName}&size=100&rounded=true&color=fff&background=random"
                    alt="avatar"
                  />
                {/if}
              </figure>
              <h4 class="fw-700 font-sm mt-2 mb-lg-5 mb-4 pl-15">
                {data?.user?.FirstName + " " + data?.user?.LastName}<span
                  class="fw-500 font-xssss text-grey-500 mt-1 mb-3 d-block"
                  >{data?.user?.Email}</span
                >
              </h4>
            </div>

            <div
              class="card-body d-block w-100 shadow-none mb-0 p-0 border-top-xs"
            >
              <ul
                class="nav nav-tabs h55 d-flex product-info-tab border-bottom-0 ps-4"
                id="pills-tab"
                role="tablist"
              >
                <li class="active list-inline-item me-5">
                  <a
                    class="fw-700 font-xssss text-grey-500 pt-3 pb-3 ls-1 d-inline-block active"
                    href="#about"
                    data-bs-toggle="pill">About</a
                  >
                </li>
                <li class="list-inline-item me-5">
                  <a
                    class="fw-700 font-xssss text-grey-500 pt-3 pb-3 ls-1 d-inline-block"
                    href="#follow"
                    data-bs-toggle="pill">follow</a
                  >
                </li>
                <li class="list-inline-item me-5">
                  <a
                    class="fw-700 font-xssss text-grey-500 pt-3 pb-3 ls-1 d-inline-block"
                    href="#following"
                    data-bs-toggle="pill">following</a
                  >
                </li>
                <!-- <li class="list-inline-item me-5"><a class="fw-700 font-xssss text-grey-500 pt-3 pb-3 ls-1 d-inline-block" href="#navtabs4" data-toggle="tab">Video</a></li> -->
                <li class="list-inline-item me-5">
                  <a
                    class="fw-700 font-xssss text-grey-500 pt-3 pb-3 ls-1 d-inline-block"
                    href="#groups"
                    data-bs-toggle="pill">Group</a
                  >
                </li>
                <li class="list-inline-item me-5">
                  <a
                    class="fw-700 font-xssss text-grey-500 pt-3 pb-3 ls-1 d-inline-block"
                    href="#events"
                    data-bs-toggle="pill">Events</a
                  >
                </li>
                <!-- <li class="list-inline-item me-5"><a class="fw-700 me-sm-5 font-xssss text-grey-500 pt-3 pb-3 ls-1 d-inline-block" href="#navtabs7" data-toggle="tab">Media</a></li> -->
              </ul>
            </div>
          </div>
        </div>
        <div class="tab-content mt-2">
          <div id="about" class="tab-pane fade show active">
            <div class="d-flex">
                <div class="col-xl-4 col-xxl-3 col-lg-4 pe-0 ">
                  <div class="card w-100 shadow-xss rounded-xxl border-0 mt-2 mb-3">
                    <div class="card-body d-block p-4">
                      <h4 class="fw-700 mb-3 font-xsss text-grey-900">About</h4>
                      <p class="fw-500 text-grey-500 lh-24 font-xssss mb-0">
                        {data?.user?.AboutMe}
                      </p>
                    </div>
    
                    <div class="card-body border-top-xs d-flex align-item-center">
                      <label class="toggle toggle-menu-color mr-2"
                        ><input
                          type="checkbox"
                          on:click={changeCountStatus}
                          disabled={toprivate}
                          bind:checked={privatestatus}
                        /><span class="toggle-icon"></span></label
                      >
                      <h4 class="fw-700 text-grey-900 font-xssss mt-2 mr-2">
                        Private
                      </h4>
                    </div>
    
                    <div class="card-body d-flex pt-0">
                      <i class="feather-map-pin text-grey-500 me-3 font-lg"></i>
                      <h4 class="fw-700 text-grey-900 font-xssss mt-1">
                        Flodia, Austia
                      </h4>
                    </div>
                    <div class="card-body d-flex pt-0">
                      <i class="feather-users text-grey-500 me-3 font-lg"></i>
                      <h4 class="fw-700 text-grey-900 font-xssss mt-1">
                        Genarel Group
                      </h4>
                    </div>
                  </div>
                  <Eventsugest/>
                </div>
                <div class="col-xl-8 col-xxl-9 col-lg-8">
                  <Createpost />
                  <Posts posts={data.createdPost} />
                </div>
            </div>
          </div>
          <div id="follow" class="tab-pane fade">
            <UserList users={data.follower}/>
          </div>
          <div id="following" class="tab-pane fade"><UserList users={data.following}/></div>
          <div id="groups" class="tab-pane fade">absabs</div>
          <div id="events" class="tab-pane fade">ashashhash</div>
        </div>
      </div>
    </div>
  </div>
</div>
<!-- main content -->

<script>
    import { onMount } from 'svelte';
    export let showAlert = false;
    export let lastMesage = {};

    function truncateMessage(message, maxLength) {
        if (message && message.length > maxLength) {
            return message.substring(0, maxLength) + '...';
        } else {
            return message;
        }
    }
</script>

<a class="alert-container" style="position: fixed; top:  15%; right:  0; z-index:  9999;">
    {#if showAlert} 
        <div class="card bg-transparent-card w-100 ps-5 mb-3 alert alert-success" style="border-left: 4px solid green; min-width: 250px;">
            <img src="/images/newmessage.png" alt="user" class="w40 position-absolute left-0" style="padding-left: 10px;">
            <h5 class="font-xsss text-grey-900 mb-1 mt-0 fw-700 d-block">{lastMesage.firstName} {lastMesage.lastName} <span class="text-grey-400 font-xsssss fw-600 float-right mt-1"></span></h5>
            <h6 class="text-grey-500 fw-500 font-xssss lh-4">{truncateMessage(lastMesage.content, 50)}</h6>
        </div>
    {/if}
</a>
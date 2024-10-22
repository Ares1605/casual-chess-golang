<script lang="ts">
  import { user } from "./lib/user";
  import type { Writable } from "svelte/store";
  import { GetOldSession } from "../wailsjs/go/main/App";
  import { AuthStatuses } from "./lib/types";
  import Loading from "./Loading.svelte";
  
  export let authStatus: Writable<AuthStatuses>;

  GetOldSession()
    .then(userResp => {
      $user = userResp;
      if ($user.setup_complete)
        $authStatus = AuthStatuses.Authenticated;
      else
        $authStatus = AuthStatuses.InitialSetup;
    })
    .catch(error => {
      console.error(error);
      $authStatus = AuthStatuses.SigningIn;
    });
</script>
<Loading />

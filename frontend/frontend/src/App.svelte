<script lang="ts">
  import { type Writable, writable } from "svelte/store";
  import { AuthStatuses } from "./lib/types";

  import SignIn from './SignInComp.svelte';
  import Authenticated from './Authenticated.svelte';
  import NotifsLayer from "./NotifsLayer.svelte";
  import ServerCheck from "./ServerCheck.svelte";
  import SetupAccount from "./SetupAccount.svelte";
  import AwaitingOldSess from "./AwaitingOldSess.svelte";

  let authStatus: Writable<AuthStatuses> = writable(AuthStatuses.AwaitingOldSess);

</script>
<NotifsLayer />
<ServerCheck>
  {#if $authStatus === AuthStatuses.AwaitingOldSess}
    <AwaitingOldSess {authStatus} />
  {:else if $authStatus === AuthStatuses.SigningIn}
    <SignIn {authStatus} />
  {:else if $authStatus === AuthStatuses.InitialSetup}
    <SetupAccount {authStatus} />
  {:else}
    <Authenticated {authStatus} />
  {/if}
</ServerCheck>

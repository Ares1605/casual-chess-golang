<script lang="ts">
  import { type Writable, writable } from "svelte/store";
  import { AuthStatuses } from "./lib/types";

  import SignIn from './pages/SignIn.svelte';
  import Authenticated from './pages/Authenticated.svelte';
  import NotifsLayer from "./NotifsLayer.svelte";
  import ServerCheck from "./pages/ServerCheck.svelte";
  import SetupAccount from "./pages/SetupAccount.svelte";
  import AwaitingOldSess from "./pages/AwaitingOldSess.svelte";
  import Modals from "./common/modals/Modals.svelte";

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
  <Modals />
</ServerCheck>

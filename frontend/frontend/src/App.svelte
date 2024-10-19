<script lang="ts">
  import SignIn from './SignInComp.svelte';
  import Home from './Home.svelte';
  import NotifsLayer from "./NotifsLayer.svelte";
  import ServerCheck from "./ServerCheck.svelte";
  import SetupAccount from "./SetupAccount.svelte";
  import { user } from "./lib/user";
  import { user as userModel } from "../wailsjs/go/models";

  enum AuthStatuses {
    Authenticated,
    SigningIn,
    InitialSetup
  };
  let authStatus = AuthStatuses.SigningIn;
  function handleSignIn(event: CustomEvent<userModel.User>) {
    $user = event.detail;
    if ($user.setup_complete)  {
      authStatus = AuthStatuses.Authenticated
    } else {
      authStatus = AuthStatuses.InitialSetup
    }
  }

</script>
<NotifsLayer />
<ServerCheck>
  {#if authStatus === AuthStatuses.SigningIn}
    <SignIn on:signIn={handleSignIn} />
  {:else if authStatus === AuthStatuses.InitialSetup}
    <SetupAccount />
  {:else}
      <Home />
  {/if}
</ServerCheck>

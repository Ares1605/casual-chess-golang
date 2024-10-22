<script lang="ts">
  import SignIn from './SignInComp.svelte';
  import Authenticated from './Authenticated.svelte';
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
  const handleSignIn = (event: CustomEvent<userModel.User>) => {
    console.log(event.detail);
    $user = event.detail;
    if ($user.setup_complete)  {
      authStatus = AuthStatuses.Authenticated;
    } else {
      authStatus = AuthStatuses.InitialSetup;
    }
  }
  const handleSetup = (() => {
    authStatus = AuthStatuses.Authenticated;
  });

</script>
<NotifsLayer />
<ServerCheck>
  {#if authStatus === AuthStatuses.SigningIn}
    <SignIn on:signIn={handleSignIn} />
  {:else if authStatus === AuthStatuses.InitialSetup}
    <SetupAccount on:setup={handleSetup} />
  {:else}
    <Authenticated />
  {/if}
</ServerCheck>

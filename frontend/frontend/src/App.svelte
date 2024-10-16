<script lang="ts">
  import type { googleuser } from "../wailsjs/go/models"
  import SignIn from './SignInComp.svelte';
  import Home from './Home.svelte';
  import AddFriend from './AddFriend.svelte';
  import NotifsLayer from "./NotifsLayer.svelte";
  import { writable, type Writable } from 'svelte/store';
  import { addingFriend } from "./lib/addingFriend"
  import { notifs, TypesType } from "./lib/notifs";

  const user: Writable<googleuser.GoogleUser | null> = writable(null);
  let isAuthenticated = false;

  function handleSignIn(event: CustomEvent<googleuser.GoogleUser>) {
    $user = event.detail;
    isAuthenticated = true;
  }

</script>
<NotifsLayer />
{#if $addingFriend}
  <AddFriend />
{/if}
{#if !isAuthenticated}
  <SignIn on:signIn={handleSignIn} />
{:else}
  <Home {user} />
{/if}

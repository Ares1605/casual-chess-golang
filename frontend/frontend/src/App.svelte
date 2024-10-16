<script lang="ts">
  import type { googleuser } from "../wailsjs/go/models"
  import SignIn from './SignInComp.svelte';
  import Home from './Home.svelte';
  import AddFriend from './AddFriend.svelte';
  import { writable, type Writable } from 'svelte/store';
  import { addingFriend } from "./lib/addingFriend.ts"

  const user: Writable<googleuser.GoogleUser | null> = writable(null);
  let isAuthenticated = false;

  function handleSignIn(event: CustomEvent<googleuser.GoogleUser>) {
    $user = event.detail;
    isAuthenticated = true;
  }
</script>
{#if $addingFriend}
  <AddFriend />
{/if}
{#if !isAuthenticated}
  <SignIn on:signIn={handleSignIn} />
{:else}
  <Home {user} {addingFriend} />
{/if}

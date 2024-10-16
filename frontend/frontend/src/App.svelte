<script lang="ts">
  import type { googleuser } from "../wailsjs/go/models"
  import SignIn from './SignInComp.svelte';
  import Home from './Home.svelte';
  import { writable, type Writable } from 'svelte/store';

  const user: Writable<googleuser.GoogleUser | null> = writable(null);
  let isAuthenticated = false;

  function handleSignIn(event: CustomEvent<googleuser.GoogleUser>) {
    $user = event.detail;
    isAuthenticated = true;
  }
</script>

{#if !isAuthenticated}
  <SignIn on:signIn={handleSignIn} />
{:else}
  <Home {user} />
{/if}

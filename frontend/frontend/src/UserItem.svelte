<script lang="ts">
  export let user: any;
  export let minimizedWidth: string;
  const src = user.profile_url;
</script>

<div class="container">
  <img {src} width={minimizedWidth} alt="profile">
  <span>{user.username}</span>
</div>

<style>
  .container {
    display: flex;
    flex-direction: row;
    justify-content: start;
    align-items: center;
    background-color: #82620a;
  }
  img {
    padding: 5px 5px;
    opacity: .9;
  }
</style>

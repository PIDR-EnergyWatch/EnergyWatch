<script>
	import { login } from '$lib/api/requests';
	import { Icon, XCircle } from 'svelte-hero-icons';
	import ulImg from '$lib/assets/Logo-universite-de-lorraine.svg_-1024x360-566710427.png';

	let username = '',
		password = '',
		error = '';

	const handleLogin = async () => {
		error = await login(username, password);
	};
</script>

<div class="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
	<div class="sm:mx-auto sm:w-full sm:max-w-sm">
		<img class="mx-auto h-20 w-auto" src={ulImg} alt="Your Company" />
		<h2 class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
			Sign in to your account
		</h2>
	</div>

	<div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
		{#if error}
			<div class="rounded-md bg-red-50 p-4">
				<div class="flex">
					<div class="flex-shrink-0">
						<Icon src={XCircle} class="h-5 w-5 text-red-400" aria-hidden="true" />
					</div>
					<div class="ml-3">
						<h3 class="text-sm font-medium text-red-800">
							{error}
						</h3>
					</div>
				</div>
			</div>
		{/if}
		<form class="space-y-6" on:submit|preventDefault={handleLogin}>
			<div>
				<label for="username" class="block text-sm font-medium leading-6 text-gray-900"
					>Username</label
				>
				<div class="mt-2">
					<input
						bind:value={username}
						id="username"
						name="username"
						type="username"
						autocomplete="username"
						required=""
						class="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
					/>
				</div>
			</div>

			<div>
				<div class="flex items-center justify-between">
					<label for="password" class="block text-sm font-medium leading-6 text-gray-900"
						>Password</label
					>
				</div>
				<div class="mt-2">
					<input
						bind:value={password}
						id="password"
						name="password"
						type="password"
						autocomplete="current-password"
						required=""
						class="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
					/>
				</div>
			</div>

			<div>
				<button
					type="submit"
					class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
					on:click={(e) => (e.target.textContent = 'Signing in...')}>Sign in</button
				>
			</div>
		</form>
	</div>
</div>

<script>
	import { logout } from '$lib/api/requests';
	import icon from '$lib/assets/ew-icon.png';
</script>

<nav class="bg-gray-800">
	<div class="px-4 sm:px-6 lg:px-8">
		<div class="flex h-16 justify-between">
			<div class="flex">
				<div class="flex flex-shrink-0 items-center">
					<img class="h-8 w-auto" src={icon} alt="Your Company" />
				</div>
				<div class="hidden md:ml-6 md:flex md:items-center md:space-x-4">
					<a
						href="/dashboard"
						class="bg-gray-900 text-white rounded-md px-3 py-2 text-sm font-medium"
						aria-current="page">Dashboard</a
					>
					<a
						href="/carousel"
						class="bg-gray-900 text-white rounded-md px-3 py-2 text-sm font-medium">Carrousel</a
					>
					<!-- <a
						href="#"
						class="text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium"
						>Projects</a
					>
					<a
						href="#"
						class="text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium"
						>Calendar</a
					> -->
				</div>
			</div>
			<div class="flex items-center">
				<div class="flex-shrink-0">
					<button
						type="button"
						class="relative inline-flex items-center gap-x-1.5 rounded-md bg-indigo-500 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500"
						on:click={async () => await logout()}
					>
						Logout
					</button>
				</div>
			</div>
		</div>
	</div>
</nav>

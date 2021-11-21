<template>
  <div>
    <Disclosure
      as="div"
      class="relative bg-sky-700 pb-32 overflow-hidden"
      v-slot="{ open }"
    >
      <div
        aria-hidden="true"
        :class="[
          open ? 'bottom-0' : 'inset-y-0',
          'absolute inset-x-0 left-1/2 transform -translate-x-1/2 w-full overflow-hidden lg:inset-y-0',
        ]"
      >
        <div class="absolute inset-0 flex">
          <div class="h-full w-1/2" style="background-color: #0a527b" />
          <div class="h-full w-1/2" style="background-color: #065d8c" />
        </div>
        <div class="relative flex justify-center">
          <svg
            class="flex-shrink-0"
            width="1750"
            height="308"
            viewBox="0 0 1750 308"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M284.161 308H1465.84L875.001 182.413 284.161 308z"
              fill="#0369a1"
            />
            <path d="M1465.84 308L16.816 0H1750v308h-284.16z" fill="#065d8c" />
            <path d="M1733.19 0L284.161 308H0V0h1733.19z" fill="#0a527b" />
            <path
              d="M875.001 182.413L1733.19 0H16.816l858.185 182.413z"
              fill="#0a4f76"
            />
          </svg>
        </div>
      </div>
      <header class="relative py-10">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <h1 class="text-3xl font-bold text-white">Programación Concurrente</h1>
        </div>
      </header>
    </Disclosure>

    <main class="relative -mt-32">
      <div class="max-w-screen-xl mx-auto pb-6 px-4 sm:px-6 lg:pb-16 lg:px-8">
        <div class="bg-white rounded-lg shadow overflow-hidden">
          <div
            class="
              divide-y divide-gray-200
              lg:grid lg:grid-cols-12 lg:divide-y-0 lg:divide-x
            "
          >
            <form
              class="divide-y divide-gray-200 lg:col-span-9"
              action="#"
              method="POST"
            >
              <!-- Profile section -->
              <div class="py-6 px-4 sm:p-6 lg:pb-8">
                <div>
                  <h2 class="text-lg leading-6 font-medium text-gray-900">
                    Información
                  </h2>
                </div>

                <div class="mt-6 flex flex-col lg:flex-row">
                  <div class="flex-grow space-y-6">
                    <div>
                      <label
                        for="about"
                        class="block text-sm font-medium text-gray-700"
                      >
                        1. Test Input
                      </label>
                      <div class="mt-1">
                        <textarea
                          id="about"
                          name="about"
                          rows="3"
                          class="
                            shadow-sm
                            focus:ring-sky-500 focus:border-sky-500
                            mt-1
                            block
                            w-full
                            sm:text-sm
                            border border-gray-300
                            rounded-md
                          "
                        />
                      </div>
                    </div>
                  </div>
                </div>

                <div class="mt-6 grid grid-cols-12 gap-6">
                  <div class="col-span-12 sm:col-span-6">
                    <label
                      for="first-name"
                      class="block text-sm font-medium text-gray-700"
                      >Test Input</label
                    >
                    <input
                      type="text"
                      name="first-name"
                      id="first-name"
                      autocomplete="given-name"
                      class="
                        mt-1
                        block
                        w-full
                        border border-gray-300
                        rounded-md
                        shadow-sm
                        py-2
                        px-3
                        focus:outline-none
                        focus:ring-sky-500
                        focus:border-sky-500
                        sm:text-sm
                      "
                    />
                  </div>

                  <div class="col-span-12 sm:col-span-6">
                    <label
                      for="last-name"
                      class="block text-sm font-medium text-gray-700"
                      >Test Input</label
                    >
                    <input
                      type="text"
                      name="last-name"
                      id="last-name"
                      autocomplete="family-name"
                      class="
                        mt-1
                        block
                        w-full
                        border border-gray-300
                        rounded-md
                        shadow-sm
                        py-2
                        px-3
                        focus:outline-none
                        focus:ring-sky-500
                        focus:border-sky-500
                        sm:text-sm
                      "
                    />
                  </div>

                  <div class="col-span-12">
                    <label
                      for="url"
                      class="block text-sm font-medium text-gray-700"
                      >Test Input</label
                    >
                    <input
                      type="text"
                      name="url"
                      id="url"
                      class="
                        mt-1
                        block
                        w-full
                        border border-gray-300
                        rounded-md
                        shadow-sm
                        py-2
                        px-3
                        focus:outline-none
                        focus:ring-sky-500
                        focus:border-sky-500
                        sm:text-sm
                      "
                    />
                  </div>

                  <div class="col-span-12 sm:col-span-6">
                    <label
                      for="company"
                      class="block text-sm font-medium text-gray-700"
                      >Test Input</label
                    >
                    <input
                      type="text"
                      name="company"
                      id="company"
                      autocomplete="organization"
                      class="
                        mt-1
                        block
                        w-full
                        border border-gray-300
                        rounded-md
                        shadow-sm
                        py-2
                        px-3
                        focus:outline-none
                        focus:ring-sky-500
                        focus:border-sky-500
                        sm:text-sm
                      "
                    />
                  </div>
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
      <div class="max-w-screen-xl mx-auto pb-6 px-4 sm:px-6 lg:pb-16 lg:px-8">
        <div class="rounded-lg shadow overflow-hidden">
          <div
            class="
              divide-y divide-gray-200
              lg:grid lg:grid-cols-12 lg:divide-y-0 lg:divide-x
            "
          >
            <form class="divide-y divide-gray-200 lg:col-span-9">
              <!-- Profile section -->
              <div class="py-6 px-4 sm:p-6 lg:pb-8">
                <div class="py-6 px-4 sm:p-6 lg:pb-8">
                  <div>
                    <h2 class="text-lg leading-6 font-medium text-gray-900">
                      Resultados
                    </h2>
                  </div>

                  <div class="mt-6 flex flex-col lg:flex-row">
                    <div class="flex-grow space-y-6">
                      <div class="flex flex-col">
                        <div class="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
                          <div
                            class="
                              py-2
                              align-middle
                              inline-block
                              min-w-full
                              sm:px-6
                              lg:px-8
                            "
                          >
                            <div
                              class="
                                shadow
                                overflow-hidden
                                border-b border-gray-200
                                sm:rounded-lg
                              "
                            >
                              <table
                                class="min-w-full divide-y divide-gray-200"
                              >
                                <thead class="bg-gray-50">
                                  <tr>
                                    <th
                                      scope="col"
                                      class="
                                        px-6
                                        py-3
                                        text-left text-xs
                                        font-medium
                                        text-gray-500
                                        uppercase
                                        tracking-wider
                                      "
                                    >
                                      Name
                                    </th>
                                    <th
                                      scope="col"
                                      class="
                                        px-6
                                        py-3
                                        text-left text-xs
                                        font-medium
                                        text-gray-500
                                        uppercase
                                        tracking-wider
                                      "
                                    >
                                      Title
                                    </th>
                                    <th
                                      scope="col"
                                      class="
                                        px-6
                                        py-3
                                        text-left text-xs
                                        font-medium
                                        text-gray-500
                                        uppercase
                                        tracking-wider
                                      "
                                    >
                                      Email
                                    </th>
                                    <th
                                      scope="col"
                                      class="
                                        px-6
                                        py-3
                                        text-left text-xs
                                        font-medium
                                        text-gray-500
                                        uppercase
                                        tracking-wider
                                      "
                                    >
                                      Role
                                    </th>
                                  </tr>
                                </thead>
                                <tbody>
                                  <tr
                                    v-for="(person, personIdx) in people"
                                    :key="person.email"
                                    :class="
                                      personIdx % 2 === 0
                                        ? 'bg-white'
                                        : 'bg-gray-50'
                                    "
                                  >
                                    <td
                                      class="
                                        px-6
                                        py-4
                                        whitespace-nowrap
                                        text-sm
                                        font-medium
                                        text-gray-900
                                      "
                                    >
                                      {{ person.name }}
                                    </td>
                                    <td
                                      class="
                                        px-6
                                        py-4
                                        whitespace-nowrap
                                        text-sm text-gray-500
                                      "
                                    >
                                      {{ person.title }}
                                    </td>
                                    <td
                                      class="
                                        px-6
                                        py-4
                                        whitespace-nowrap
                                        text-sm text-gray-500
                                      "
                                    >
                                      {{ person.email }}
                                    </td>
                                    <td
                                      class="
                                        px-6
                                        py-4
                                        whitespace-nowrap
                                        text-sm text-gray-500
                                      "
                                    >
                                      {{ person.role }}
                                    </td>
                                  </tr>
                                </tbody>
                              </table>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import { ref } from "vue";
import {
  Disclosure,
  DisclosureButton,
  DisclosurePanel,
  Menu,
  MenuButton,
  MenuItem,
  MenuItems,
  Switch,
  SwitchDescription,
  SwitchGroup,
  SwitchLabel,
} from "@headlessui/vue";
import { SearchIcon } from "@heroicons/vue/solid";
import {
  BellIcon,
  CogIcon,
  CreditCardIcon,
  KeyIcon,
  MenuIcon,
  UserCircleIcon,
  ViewGridAddIcon,
  XIcon,
} from "@heroicons/vue/outline";

const people = [
  {
    name: "Resultado 1",
    title: "Resultado 1",
    role: "Resultado 1",
    email: "Resultado 1",
  },
  {
    name: "Resultado 2",
    title: "Resultado 2",
    role: "Resultado 2",
    email: "Resultado 2",
  },
  // More people...
];

const user = {
  name: "Debbie Lewis",
  handle: "deblewis",
  email: "debbielewis@example.com",
  imageUrl:
    "https://images.unsplash.com/photo-1517365830460-955ce3ccd263?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=320&h=320&q=80",
};
const navigation = [
  { name: "Dashboard", href: "#", current: true },
  { name: "Jobs", href: "#", current: false },
  { name: "Applicants", href: "#", current: false },
  { name: "Company", href: "#", current: false },
];
const subNavigation = [
  { name: "Profile", href: "#", icon: UserCircleIcon, current: true },
  { name: "Account", href: "#", icon: CogIcon, current: false },
  { name: "Password", href: "#", icon: KeyIcon, current: false },
  { name: "Notifications", href: "#", icon: BellIcon, current: false },
  { name: "Billing", href: "#", icon: CreditCardIcon, current: false },
  { name: "Integrations", href: "#", icon: ViewGridAddIcon, current: false },
];
const userNavigation = [
  { name: "Your Profile", href: "#" },
  { name: "Settings", href: "#" },
  { name: "Sign out", href: "#" },
];

export default {
  components: {
    Disclosure,
    DisclosureButton,
    DisclosurePanel,
    Menu,
    MenuButton,
    MenuItem,
    MenuItems,
    Switch,
    SwitchDescription,
    SwitchGroup,
    SwitchLabel,
    BellIcon,
    MenuIcon,
    SearchIcon,
    XIcon,
  },
  setup() {
    const availableToHire = ref(true);
    const privateAccount = ref(false);
    const allowCommenting = ref(true);
    const allowMentions = ref(true);

    return {
      user,
      navigation,
      subNavigation,
      userNavigation,
      availableToHire,
      privateAccount,
      allowCommenting,
      allowMentions,
      people,
    };
  },
};
</script>
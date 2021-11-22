<template>
<div>
    <div class="relative min-h-screen  grid bg-black ">
      <div class="flex flex-col sm:flex-row items-center md:items-start sm:justify-center md:justify-start flex-auto min-w-0 ">
        <div  class="relative sm:w-1/2 xl:w-3/5 bg-blue-500 h-full hidden md:flex flex-auto items-center justify-center p-10 overflow-hidden  text-white bg-no-repeat bg-cover relative" style="background-image: url(https://s03.s3c.es/imag/_v0/770x420/2/c/c/UCP-Monterrico-3.jpg);">
          <div class="absolute bg-black  opacity-25 inset-0 z-0"></div>
          <div class="w-full  lg:max-w-2xl md:max-w-md z-10 items-center text-center ">
            <div class=" font-bold leading-tight mb-6 mx-auto w-full content-center items-center ">
          
            </div>
          </div>
        </div>

        <div class="md:flex md:items-center md:justify-left w-full sm:w-auto md:h-full xl:w-1/2 p-8  md:p-10 lg:p-14 sm:rounded-lg md:rounded-none ">
            <div class="max-w-xl w-full space-y-12">
              <div class="lg:text-left text-center">
          
                <div class="flex items-center justify-center ">
                  <div class="bg-black flex flex-col w-80 border border-white rounded-lg px-8 py-10">
                  <form action="/" method="post">
            <div class="flex flex-col mb-4">
                <label class="mb-2 font-bold text-sm text-white" for="first_name">Monóxido de Carbono (ug/m3)</label>
                <input class="border py-2 px-3 text-grey-800" type="text" name="first_name" id="first_name">
            </div>
            <div class="flex flex-col mb-4">
                <label class="mb-2 font-bold text-sm text-white" for="last_name">Ácido Sulfhídrico (ug/m3)</label>
                <input class="border py-2 px-3 text-grey-800" type="text" name="last_name" id="last_name">
            </div>
            <div class="flex flex-col mb-4">
                <label class="mb-2 font-bold text-sm text-white" for="last_name">Dióxido de Nitrógeno (ug/m3)</label>
                <input class="border py-2 px-3 text-grey-800" type="text" name="last_name" id="last_name">
            </div>
            <div class="flex flex-col mb-4">
                <label class="mb-2 font-bold text-sm text-white" for="last_name">Ozono (ug/m3)</label>
                <input class="border py-2 px-3 text-grey-800" type="text" name="last_name" id="last_name">
            </div>
            <div class="flex flex-col mb-4">
                <label class="mb-2 font-bold text-sm text-white" for="last_name">PM10 (ug/m3)</label>
                <input class="border py-2 px-3 text-grey-800" type="text" name="last_name" id="last_name">
            </div>
            <div class="flex flex-col mb-4">
                <label class="mb-2 font-bold text-sm text-white" for="last_name">PM2,5 (ug/m3)</label>
                <input class="border py-2 px-3 text-grey-800" type="text" name="last_name" id="last_name">
            </div>
            <div class="flex flex-col mb-4">
                <label class="mb-2 font-bold text-sm text-white" for="last_name">Dióxido de Azufre(ug/m3)</label>
                <input class="border py-2 px-3 text-grey-800" type="text" name="last_name" id="last_name">
            </div>
            <div class="flex flex-col mb-4">
                <label class="mb-2 font-bold text-sm text-white" for="last_name">Ruido (dB)</label>
                <input class="border py-2 px-3 text-grey-800" type="text" name="last_name" id="last_name">
            </div>
            <div class="flex flex-col mb-4">
                <label class="mb-2 font-bold text-sm text-white" for="last_name">UV</label>
                <input class="border py-2 px-3 text-grey-800" type="text" name="last_name" id="last_name">
            </div>
            <div class="flex flex-col mb-4">
                <label class="mb-2 font-bold text-sm text-white" for="last_name">Humedad (%)</label>
                <input class="border py-2 px-3 text-grey-800" type="text" name="last_name" id="last_name">
            </div>
            <div class="flex flex-col mb-4">
                <label class="mb-2 font-bold text-sm text-white" for="last_name">Presión (hPa)</label>
                <input class="border py-2 px-3 text-grey-800" type="text" name="last_name" id="last_name">
            </div>
            
            <button class="block bg-green-400 hover:bg-green-600 text-white uppercase text-lg mx-auto p-4 rounded" type="submit">Enviar</button>
        </form>
          
                </div>
                </div>
              
            </div>
              
            </div>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
import { ref } from "vue";

export default {
  
};
</script>
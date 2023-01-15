@extends('dashboard.layouts.main')

@section('container')
<div class="row">
  <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <h1 class="h2">Selamat Datang {{ auth()->user()->nama_lengkap }}</h1>
      {{-- <div class="btn-toolbar mb-2 mb-md-0">
        <div class="btn-group me-2">
          <button type="button" class="btn btn-sm btn-outline-secondary">Share</button>
          <button type="button" class="btn btn-sm btn-outline-secondary">Export</button>
        </div>
        <button type="button" class="btn btn-sm btn-outline-secondary dropdown-toggle">
          <span data-feather="calendar"></span>
          This week
        </button>
      </div> --}}
    </div>
    <div class="row my-3">
      <div class="col-md-3">
        <a class="btn d-block border-0 text-decoration-none linka" href="#" style="background-color: rgba(104, 91, 19, 0.782) ">Buat Catatan</a>
      </div>
    </div>
    <div class="row">
      <div class="col-md-4 mb-3">
        <div class="card shadow mb-2 bg-body rounded">
          <button type="button" class="btn-close ms-auto py-2 px-2" aria-label="Close"></button>
          <div class="card-body">
            <h5 class="card-title">Catatan Pertama</h5>
            <p class="card-text">Some quick example text to build on the hkjahjh;ahfajlfjdkdjsad aljfakcard title and make up the bulk of the card's content.</p>
            <a href="#" class="btn ms-auto py-2 px-2e"><i class="bi bi-zoom-in"></i></a>
          </div>
        </div>
      </div>
      <div class="col-md-4 mb-3">
        <div class="card shadow mb-2 bg-body rounded">
          <button type="button" class="btn-close ms-auto py-2 px-2" aria-label="Close"></button>
          <div class="card-body">
            <h5 class="card-title">Catatan Kedua </h5>
            <p class="card-text">Some quick example text to build on the hkjahjh;ahfajlfjdkdjsad aljfakcard title and make up the bulk of the card's content.</p>
            <a href="#" class="btn ms-auto py-2 px-2e"><i class="bi bi-zoom-in"></i></a>
  
          </div>
        </div>
      </div>
      <div class="col-md-4 mb-3">
        <div class="card shadow mb-2 bg-body rounded">
          <button type="button" class="btn-close ms-auto py-2 px-2" aria-label="Close"></button>
          <div class="card-body">
            <h5 class="card-title">Catatan Ketiga</h5>
            <p class="card-text">Some quick example text to build on the hkjahjh;ahfajlfjdkdjsad aljfakcard title and make up the bulk of the card's content.</p>
            <a href="#" class="btn ms-auto py-2 px-2e"><i class="bi bi-zoom-in"></i></a>
  
          </div>
        </div>
      </div>
      <div class="col-md-4 mb-3">
        <div class="card shadow mb-2 bg-body rounded">
          <button type="button" class="btn-close ms-auto py-2 px-2" aria-label="Close"></button>
          <div class="card-body">
            <h5 class="card-title">Catatan Keempat</h5>
            <p class="card-text">Some quick example text to build on the hkjahjh;ahfajlfjdkdjsad aljfakcard title and make up the bulk of the card's content.</p>
            <a href="#" class="btn ms-auto py-2 px-2e"><i class="bi bi-zoom-in"></i></a>
  
          </div>
        </div>
      </div>
      <div class="col-md-4 mb-3">
        <div class="card shadow mb-2 bg-body rounded">
          <button type="button" class="btn-close ms-auto py-2 px-2" aria-label="Close"></button>
          <div class="card-body">
            <h5 class="card-title">Catatan Kelima</h5>
            <p class="card-text">Some quick example text to build on the hkjahjh;ahfajlfjdkdjsad aljfakcard title and make up the bulk of the card's content.</p>
            <a href="#" class="btn ms-auto py-2 px-2e"><i class="bi bi-zoom-in"></i></a>
  
          </div>
        </div>
      </div>
    </div>

</div>

    
@endsection

@push('script')
<script type="text/javascript">

window.addEventListener('keydown',function(e){
        if(e.ctrlKey && (e.which==83)){
          alert('@copyright RogerDev')
            e.preventDefault();
        }
    });

    function showTime() {
     var a_p = "WIB";
     var today = new Date();
     var curr_hour = today.getHours();
     var curr_minute = today.getMinutes();
     var curr_second = today.getSeconds();
     
     if (curr_hour > 24) {
         curr_hour = curr_hour - 24;
     }
     curr_hour = checkTime(curr_hour);
     curr_minute = checkTime(curr_minute);
     curr_second = checkTime(curr_second);
 document.getElementById('clock').innerHTML="Pukul: " + curr_hour + ":" + curr_minute + ":" + curr_second + " " + a_p;
     }
 
 function checkTime(i) {
     if (i < 10) {
         i = "0" + i;
     }
     return i;
 }
 setInterval(showTime, 500);


 
// Menampilkan Hari, Bulan dan Tahun 
 
 var months = ['Januari', 'Februari', 'Maret', 'April', 'Mei', 'Juni', 'Juli', 'Agustus', 'September', 'Oktober', 'November', 'Desember'];
 var myDays = ['Minggu', 'Senin', 'Selasa', 'Rabu', 'Kamis', 'Jum&#39;at', 'Sabtu'];
 var date = new Date();
 var day = date.getDate();
 var month = date.getMonth();
 var thisDay = date.getDay(),
     thisDay = myDays[thisDay];
 var yy = date.getYear();
 var year = (yy < 1000) ? yy + 1900 : yy;
 document.getElementById('tanggal').innerHTML= thisDay + ', ' + day + ' ' + months[month] + ' ' + year;

</script>
    
@endpush
<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="csrf-token" content="{{ csrf_token() }}">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="">
    <meta name="author" content="RogerDev">
    <meta name="generator" content="Hugo 0.87.0">
    <title>Dashboard </title>
    

    <link rel="stylesheet" href="https://cdn.datatables.net/1.11.3/css/jquery.dataTables.min.css">
    <link rel="stylesheet" href="https://cdn.datatables.net/buttons/2.0.1/css/buttons.dataTables.min.css">
    <!-- Bootstrap core CSS -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-1BmE4kWBq78iYhFldvKuhfTAU6auU8tT94WrHftjDbrCEXSU1oBoqyl2QvZ6jIW3" crossorigin="anonymous">
    {{-- datepicker --}}
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-datepicker/1.9.0/css/bootstrap-datepicker.min.css" integrity="sha512-mSYUmp1HYZDFaVKK//63EcZq4iFWFjxSL+Z3T/aCt4IO9Cejm03q3NKKYN6pFQzY0SBOr8h+eCIAZHPXcpZaNw==" crossorigin="anonymous" referrerpolicy="no-referrer" />

    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.7.0/font/bootstrap-icons.css">
    

    
    {{-- select2 --}}
    <link href="https://cdn.jsdelivr.net/npm/select2@4.1.0-rc.0/dist/css/select2.min.css" rel="stylesheet" />


    <!-- Custom styles for this template -->
    <link href="/css/dashboard.css" rel="stylesheet">
    <link rel="stylesheet" href="/css/style.css">
    
    
  </head>
  <body>

    @include('dashboard.layouts.header')

<div class="container-fluid">
  <div class="row">
    <div class="col-md-3 col-lg-2 collapse show" id="sidebarCollapse">
      
      @include('dashboard.layouts.sidebar')
    </div>
    <div class="" id="sembunyikanSidebar">
      <h1>
        <i class=" fixed-top bi bi-arrow-right-circle-fill" style="top: 50%; width:30px;" hidden></i>
      </h1>
    </div>
    <main class="col-md-9 ms-sm-auto col-lg-10 px-md-4" id="main">
     @yield('container')

     {{-- modal sidebar logout --}}
     <div class="modal fade" id="modalLogout" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel">KELUAR AKUN</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            Yakin untuk keluar?
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
            <form action="/logout" method="POST">        
              @csrf    
            <button type="submit" class=" btn btn-primary">Logout</button>
          </form>
          </div>
        </div>
      </div>
    </div>
    </main>
  </div>
</div>

<script src="https://code.jquery.com/jquery-3.5.1.js"></script>

    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-ka7Sk0Gln4gmtz2MlQnikT1wXgYsOg+OMhuP+IlRH9sENBO0LRn5q+8nbTov4+1p" crossorigin="anonymous"></script>

      <script src="https://cdn.jsdelivr.net/npm/feather-icons@4.28.0/dist/feather.min.js" integrity="sha384-uO3SXW5IuS1ZpFPKugNNWqTZRRglnUJK6UAZ/gxOX80nxEkN9NcGZTftn6RzhGWE" crossorigin="anonymous"></script>
      <script src="https://cdn.jsdelivr.net/npm/chart.js@2.9.4/dist/Chart.min.js" integrity="sha384-zNy6FEbO50N+Cg5wap8IKA4M/ZnLJgzc6w2NqACZaK0u0FXfOWRRJOnQtpZun8ha" crossorigin="anonymous"></script>
      <script type="text/javascript" src="https://cdn.datatables.net/v/bs4/dt-1.11.3/b-2.0.1/datatables.min.js"></script>
      <script src="https://cdn.datatables.net/1.11.3/js/jquery.dataTables.min.js"></script>
      <script src="https://cdn.datatables.net/buttons/2.0.1/js/dataTables.buttons.min.js"></script>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/jszip/3.1.3/jszip.min.js"></script>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/pdfmake/0.1.53/pdfmake.min.js"></script>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/pdfmake/0.1.53/vfs_fonts.js"></script>
      <script src="https://cdn.datatables.net/buttons/2.0.1/js/buttons.html5.min.js"></script>
      <script src="https://cdn.datatables.net/buttons/2.0.1/js/buttons.print.min.js"></script>

{{-- CDN untuk edit tampilan di exel  --}}


<script src="https://cdn.jsdelivr.net/npm/datatables-buttons-excel-styles@1.2.0/js/buttons.html5.styles.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/datatables-buttons-excel-styles@1.2.0/js/buttons.html5.styles.templates.min.js"></script>
      {{-- <script type="text/javascript" src="/js/dashboard.js"></script> --}}

      {{-- datepicker --}}
<script src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-datepicker/1.9.0/js/bootstrap-datepicker.min.js" integrity="sha512-T/tUfKSV1bihCnd+MxKD0Hm1uBBroVYBOYSk1knyvQ9VyZJpc/ALb4P0r6ubwVPSGB2GvjeoMAJJImBG12TiaQ==" crossorigin="anonymous" referrerpolicy="no-referrer"></script>
      @stack('script')

       <script>
        $(function () {
          $('.bi-arrow-right-circle-fill').click( function(e){
            e.preventDefault();
            $('#sidebarCollapse').removeAttr('hidden');
            $(this).attr('hidden', 'true');
            $('#main').attr('class', 'col-md-9 ms-sm-auto col-lg-10 px-md-4');
            $('#tombolRefreshTable').trigger('click');

          });

          $('#sidebarCollapseToggle').click( function(e){
            e.preventDefault();
            
            $('.bi-arrow-right-circle-fill').removeAttr('hidden');
            $('#sidebarCollapse').attr('hidden', 'true');
            $('#tombolRefreshTable').trigger('click');
            $('#main').attr('class', 'col-md-12 ms-sm-auto col-lg-12 px-md-4');
            


          });

        });

        //UNTUK MWNONAKTIVEKAN KLIK KANAN

      //   $(document).bind("contextmenu",function(e) {
      //   alert('Tidak diperbolehkan');//silahkan tulis pesan yang akan ditampilkan
      //   return false;
      // });

      //   document.onkeydown = function(e) {
      //   if(event.keyCode == 123) {
      //     return false;
      //     }
      //     if(e.ctrlKey && e.shiftKey && e.keyCode == 'I'.charCodeAt(0)){
      //     alert('Tidak diperbolehkan')
      //     return false;
      //     }
      //     if(e.ctrlKey && e.shiftKey && e.keyCode == 'J'.charCodeAt(0)){
      //     alert('Tidak diperbolehkan')
      //     return false;
      //     }
      //     if(e.ctrlKey && e.keyCode == 'U'.charCodeAt(0)){
      //     alert('Tidak diperbolehkan')
      //     return false;
      //     }
      //   }

      </script>
  </body>
</html>

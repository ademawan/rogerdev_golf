<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
    <div class="position-sticky pt-3">
      <ul class="nav flex-column">
        <li class="nav-item">
            <a class="nav-link" href="/dashboard">
              <span data-feather="home"></span>DASHBOARD
            </a>
          </li>

          <li><hr class="dropdown-divider"></li>

        <li class="nav-item">
            <a class="nav-link" data-bs-toggle="collapse" href="#collapseExample" role="button" aria-expanded="false" aria-controls="collapseExample">
              <span data-feather="home"></span> MENU MASTER
            </a>
          </li>
          <div class="collapse {{ Request::is('dashboard/master*') ? 'show' : '' }}" id="collapseExample">
            <div class="bg-light collapse-inner rounded">
                    <li><a class="dropdown-item {{ Request::is('dashboard/master/daftar-alat-kerja') ? 'active' : '' }}" href="/dashboard/master/daftar-alat-kerja">Daftar Alat Kerja</a></li>
                    <li><a class="dropdown-item {{ Request::is('dashboard/master/daftar-mesin') ? 'active' : '' }}" href="/dashboard/master/daftar-mesin">Daftar Mesin</a></li>
                    <li><a class="dropdown-item {{ Request::is('dashboard/master/daftar-sparepart') ? 'active' : '' }}" href="/dashboard/master/daftar-sparepart">Daftar Sparepart</a></li>
                    <li><a class="dropdown-item {{ Request::is('dashboard/master/daftar-teknisi') ? 'active' : '' }}" href="/dashboard/master/daftar-teknisi">Daftar Teknisi</a></li>
                    <li><a class="dropdown-item {{ Request::is('dashboard/master/daftar-user') ? 'active' : '' }}" href="/dashboard/master/daftar-user">Daftar User</a></li>
            </div>
          </div>

        <li><hr class="dropdown-divider"></li>

        <li class="nav-item">
          <a class="nav-link" data-bs-toggle="collapse" href="#colallapseForm" role="button" aria-expanded="false" aria-controls="collapseExample">
            <span data-feather="home"></span> MENU FORM
          </a>
        </li>
        <div class="collapse {{ Request::is('dashboard/form*') ? 'show' : '' }}" id="colallapseForm">
          <div class="bg-light collapse-inner rounded">
                  <li><a class="dropdown-item {{ Request::is('dashboard/form/absensi-teknisi-maintenance') ? 'active' : '' }}" href="/dashboard/form/absensi-teknisi-maintenance">Absensi Teknisi Maintenance</a></li>
                  <li><a class="dropdown-item {{ Request::is('dashboard/form/permintaan-perbaikan') ? 'active' : '' }}" href="/dashboard/form/permintaan-perbaikan">Permintaan Perbaikan</a></li>
                  <li><a class="dropdown-item {{ Request::is('dashboard/form/penambahan-stock-sparepart') ? 'active' : '' }}" href="/dashboard/form/penambahan-stock-sparepart">Penambahan Stock Sparepart</a></li>
          </div>
        </div>

        <li><hr class="dropdown-divider"></li>

        <li class="nav-item">
          <a class="nav-link" data-bs-toggle="collapse" href="#collapseLaporan" role="button" aria-expanded="false" aria-controls="collapseExample">
            <span data-feather="home"></span> MENU LAPORAN
          </a>
        </li>
        <div class="collapse {{ Request::is('dashboard/laporan*') ? 'show' : '' }}" id="collapseLaporan">
          <div class="bg-light collapse-inner rounded">
                  <li><a class="dropdown-item {{ Request::is('dashboard/laporan/laporan-absensi') ? 'active' : '' }}" href="/dashboard/laporan/laporan-absensi">Laporan Absensi</a></li>
                  <li><a class="dropdown-item {{ Request::is('dashboard/laporan/laporan-pekerjaan') ? 'active' : '' }}" href="/dashboard/laporan/laporan-pekerjaan">Laporan Pekerjaan</a></li>
                  <li><a class="dropdown-item {{ Request::is('dashboard/laporan/laporan-masuk-dan-keluar') ? 'active' : '' }}" href="/dashboard/laporan/laporan-masuk-dan-keluar">Laporan Masuk dan keluar</a></li>
          </div>
        </div>
        <li><hr class="dropdown-divider"></li>

        <li class="nav-item">
          <button type="button" class="nav-link border-0 bg-transparent text-dark" data-bs-toggle="modal" data-bs-target="#modalLogout">
            KELUAR AKUN
          </button>
        </li>
        <li><hr class="dropdown-divider"></li>

        <li class="nav-item">
          <a class="nav-link" data-bs-toggle="collapse" id="sidebarCollapseToggle" role="button" aria-expanded="false" aria-controls="collapseExample">
            <span data-feather="home"></span>
            <h1 class="text-center">
              <i class="bi bi-arrow-left-circle-fill"></i>
            </h1>
              

          </a>
        </li>

        


      </ul>
      <ul class="nav d-block flex-column fixed-bottom w-25 responstime">
        <li class="nav-item ">
          <a class="nav-link" data-bs-toggle="collapse" href="#" role="button" aria-expanded="false" aria-controls="collapseExample">
            <span id="tanggal" style="background-color: rgb(235, 159, 17)"></span>
            <H5 id="clock"></H5>
          </a>
        </li>
      </ul>
    </div>


    {{-- modal Logout --}}
    <!-- Button trigger modal -->





  </nav>
@extends('dashboard.layouts.main')

@section('container')
<div class="row">
<div class="d-flex justify-content-center flex-wrap flex-md-nowrap align-items-center border-bottom">
    <h1 class="h2 d-blockr">DAFTAR USER</h1>




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

  {{-- <div class="row">
    {{-- <div class="col">

      <button type="submit">Excel</button>
      <button type="submit">PDF</button>
      <button type="submit">Print</button>
    </div>
    <div class="col">

      <button class="float-end" type="submit">Search</button><input class="float-end" type="text " placeholder="Search....">
     
    </div>
  </div> --}}
  <div class="row my-3">
    <div class="col-md-12">
      <button type="button" id="tombolTambahUser" class="btn linka d-inline mb-2" style="background-color: rgba(104, 91, 19, 0.782) " data-bs-toggle="collapse" data-bs-target="#tambahuser" aria-expanded="false" aria-controls="collapseExample">
        Tambah User
      </button>
      <div hidden>
        <button type="button" id="tombolRefreshTable" class="btn linka d-inline mb-2" style="background-color: rgba(104, 91, 19, 0.782) " >
          Refresh Table
        </button>

      </div>
    </div>
      
    <div class="col-md-12 collapse" id="tambahuser">
      <div class="col-md">
        <form class="form-horizontal form-disable-enter" name="formtambahuser" id="formtambahuser">
          @csrf
          <input type="hidden" id="id_user" name="id_user" class="form-control" value="">
          <div class="row justify-content-around">
       
            <div class="col-lg-5 shadow-sm">
            
              <div data-spy="scroll" id="scroll-ku" data-bs-target="#scroll-ku" data-bs-offset="0" style="height: 300px;overflow: auto;">
                <div class="row py-2" style="margin:0 !important;">

                  <div class="col-md-12" >
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Nama Lengkap</label>
                      <div class="col-sm-9">
                        <input type="text" id="nama_lengkap" name="nama_lengkap" class="form-control" autofocus>
                        <span class="text-danger error-text nama_lengkap_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Username</label>
                      <div class="col-sm-9">
                        <input type="text" id="username" name="username" class="form-control">
                        <span class="text-danger error-text username_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Email</label>
                      <div class="col-sm-9">
                        <input type="text" id="email" name="email" class="form-control">
                        <span class="text-danger error-text email_error"></span>
                      </div>
                    </div>
                  </div>
                  
                  <div class="col-md-12" id="colpassword">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Password</label>
                      <div class="col-sm-9">
                        <input type="text" id="password" name="password" class="form-control">
                        <span class="text-danger error-text password_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Jabatan</label>
                      <div class="col-sm-9">
                        <input type="text" id="jabatan" name="jabatan" class="form-control">
                         <span class="text-danger error-text jabatan_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Status</label>
                      <div class="col-sm-9">
                        <select class="form-select" id="status" name="status" aria-label="Default select example">
                          <option value="1" selected>Aktif</option>
                          <option value="0" >Tidak Aktif</option>
                        </select>
                      </div>
                    </div>      
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Role</label>
                      <div class="col-sm-9">
                        <select class="form-select" id="role" name="role" aria-label="Default select example">
                          <option value="1" selected>Admin Maintenance</option>
                          <option value="2" >HRD</option>
                        </select>
                      </div>
                    </div>      
                  </div>
                  
                </div>
                <div class="d-grid gap-2 d-md-flex justify-content-md-center telusuri">
                    <button id="saveBtnUpdateUser" class="btn btn-primary me-md-2" type="submit">Tambah</button>
                    <button id="saveBtnTambahUser" class="btn btn-primary me-md-2" type="submit">Tambah</button>
                    <button id="saveBtnBatalUser" class="btn btn-danger updateBtnBatalUser" type="submit">Batal</button>
                </div>
       
              </div>
            </div>
            <div class="col-lg-6 shadow-sm">Lorem ipsum dolor sit amet consectetur adipisicing elit. Hic maxime quidem modi repellat fugit velit explicabo nostrum sequi, blanditiis eum voluptas deserunt culpa nemo ut recusandae at nulla illo officiis dolores, cupiditate deleniti vel libero. Accusantium fugit officiis repellat eos sed, voluptatem facere reiciendis eaque inventore cupiditate repellendus veniam veritatis corrupti dignissimos voluptas. Eveniet obcaecati iste dignissimos autem.
              <div class="row">
                <small>
                <span class="text-danger error-text nama_lengkap_errornotif"></span>
                <span class="text-danger error-text username_errornotif"></span>
                <span class="text-danger error-text email_errornotif"></span>
                <span class="text-danger error-text password_errornotif"></span>
                <span class="text-danger error-text jabatan_errornotif"></span>
                <span class="text-danger error-text status_errornotif"></span>
                <span class="text-danger error-text role_errornotif"></span>
                </small>
              </div>
            </div>
            
          </div>

          </div>

      </form>

      </div>
    </div>
    
  </div>
  <hr class="dropdown-divider mb-3">
 
    <div class="table-responsive">
      <table id="myTable" class="table table-striped" width="100%">
        <thead>
          <tr>
            <th>Nama Lengkap</th>
            <th>Username</th>
            <th>Email</th>  
            <th>Jabatan</th>  
            <th>Status</th>  
            <th>Role</th>  
            <th>Action</th>
          </tr>     
        </thead>
        <tbody>
          
        </tbody>
      </table>
    </div>


</div>    
@endsection

@push('script')
<script src="/js/master/user.js"></script>
    
@endpush
import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit{
	myForm = new FormGroup({
    		name: new FormControl('', [Validators.required, Validators.minLength(3)]),
    		file: new FormControl('', [Validators.required]),
    		fileSource: new FormControl('', [Validators.required])
  	});

	constructor(private http: HttpClient) {}
  	
	ngOnInit(): void {
	}

	get f(){
    		return this.myForm.controls;
  	}

	onFileChange(event:any) { 
    		if (event.target.files.length > 0) {
      			const file = event.target.files[0];
      			this.myForm.patchValue({
        			fileSource: file
      			});
    		}
  	}

   	submit(){
    		const formData = new FormData();
    		formData.append('file', this.myForm.get('fileSource')?.value);
    		this.http.post('http://localhost:8080/upload', formData)
    			.subscribe(res => console.log(res));
   	}
}

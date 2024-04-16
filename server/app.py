from flask import Flask, redirect, request, send_file
from flask_pymongo import PyMongo
import datetime, pytz, io, csv





app = Flask(__name__)
app.config["MONGO_URI"] = "" # enter your mongodb db uri to save data
app.secret_key = "" # enter your server encryption key

mongo_client = PyMongo(app)
db = mongo_client.db


@app.route('/')
def hello():
    return redirect("https://youtu.be/gdEe_lXsULg") # trolling

@app.route("/api/status", methods=["GET"])
def api_status():
    try:
        return {"status":"good", "msg" : "Server is alive"}, 200
    except:
        return "error", 500

@app.route("/api/register", methods=["POST"])
def register_user():
    try:
        request_json = request.get_json()
        one_bil = request_json["one_bil"]
        hundred_mil = request_json["hundred_mil"]
        ten_mil = request_json["ten_mil"]
        one_mil = request_json["one_mil"]
        cpu = request_json["cpu"]
        
        try:
            new_user = db.results.insert_one({
                                            "one_bil" : one_bil,
                                            "hundred_mil" : hundred_mil,
                                            "ten_mil" : ten_mil,
                                            "one_mil" : one_mil,
                                            "cpu" : cpu,
                                            "registration_date" : datetime.datetime.now(pytz.timezone("Europe/Riga")),
                                            })
            
            print("results aquired succesfully cpu - " + cpu)
            return {"status":"good", "msg" : "Results registered successfully!"}, 200

        except:
            return {"status":"bad", "msg" : "Error while registering results. Please contact admins."}, 500
        
        


    except:
        return "error", 500


@app.route("/api/get_results_csv")
def get_results_csv_view():
    try:

        data = db.results.find({})
        csvfile = io.BytesIO()
        for entry in data:
            one_bil = entry["one_bil"]
            hundred_mil = entry["hundred_mil"]
            ten_mil = entry["ten_mil"]
            one_mil = entry["one_mil"]
            cpu = entry["cpu"]
            row = f"{cpu},{one_bil},{hundred_mil},{ten_mil},{one_mil}\n".encode("utf-8")
            csvfile.write(row)
        csvfile.seek(0)

        print(1)
        return send_file(csvfile, download_name="data.csv", as_attachment=True)
    except:
        return "error", 500

if __name__ == '__main__':
    if app.config["MONGO_URI"] or "" or app.secret_key == "":
        print("please setup server variables at top of app.py")
    else:
        app.run()
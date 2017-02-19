from sanic import Sanic
from sanic.response import redirect, text
import boto3
import botocore

app = Sanic()
s3 = boto3.client('s3')


@app.route("/")
async def no_file(request):
    '''
        No file specified
    '''
    return text("File Does Not Exist")


@app.route("/<name>")
async def file(request, name):
    '''
        Handler for sending link for downloading file
    '''
    file_name = name
    bucket = "calpolyfast-ctf-winter17"

    # Check if bucket name given
    try:
        bucket_temp = request.args["bucket"][0]
        if bucket is not None:
            bucket = bucket_temp
    except:
        pass

    # Check if folder given
    try:
        folder = request.args["folder"][0]
        if folder is not None:
            if folder[len(folder)-1:] == "/":
                file_name = folder + file_name
            else:
                file_name = folder + "/" + file_name
    except:
        pass

    # Check if bucket exists
    try:
        s3.head_bucket(
            Bucket=bucket,
        )
    except botocore.exceptions.ClientError:
        return text("Bucket Does Not Exist")

    # Checks if file exists
    try:
        s3.head_object(
            Bucket=bucket,
            Key=file_name
        )
    except botocore.exceptions.ClientError:
        return text("File Does Not Exist")

    # Get temporary url for file that expires
    url = s3.generate_presigned_url(
        "get_object",
        Params={
            "Bucket": bucket,
            "Key": file_name
        }
    )

    return redirect(url)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8000)

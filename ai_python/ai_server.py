from fastapi import FastAPI
from pydantic import BaseModel
import tensorflow as tf
import numpy as np

app = FastAPI()

# Load TensorFlow Model
model = tf.keras.models.load_model("task_priority_model.h5")

# Define request body model
class TaskInput(BaseModel):
    priority_level: int
    days_until_deadline: float

@app.post("/predict/")
def predict(data: TaskInput):
    # Prepare input for the model
    input_data = np.array([[data.priority_level, data.days_until_deadline]])
    
    # Make prediction
    prediction = model.predict(input_data)
    predicted_priority = int(np.argmax(prediction))

    return {"predicted_priority": predicted_priority}

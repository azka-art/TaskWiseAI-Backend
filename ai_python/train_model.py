import tensorflow as tf
import numpy as np

# Sample training data: [priority_level, days_until_deadline]
X_train = np.array([
    [0, 10],  # Low priority, 10 days left
    [1, 5],   # Medium priority, 5 days left
    [2, 2],   # High priority, 2 days left
    [2, 1],   # High priority, 1 day left
    [0, 15],  # Low priority, 15 days left
    [1, 7]    # Medium priority, 7 days left
])
y_train = np.array([0, 1, 2, 2, 0, 1])  # Labels: Low, Medium, High

# Define a simple neural network model
model = tf.keras.Sequential([
    tf.keras.layers.Dense(10, activation='relu', input_shape=(2,)),
    tf.keras.layers.Dense(3, activation='softmax')  # 3 classes: Low, Medium, High
])

# Compile the model
model.compile(optimizer='adam', loss='sparse_categorical_crossentropy', metrics=['accuracy'])

# Train the model
model.fit(X_train, y_train, epochs=50)

# Save the model as 'task_priority_model.h5'
model.save("task_priority_model.h5")

print("✅ Model trained and saved as 'task_priority_model.h5'!")

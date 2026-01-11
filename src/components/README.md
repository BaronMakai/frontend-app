import subprocess
import os

def build_frontend():
    """Builds the frontend application."""
    try:
        # Navigate to the frontend directory
        os.chdir('./frontend')
        
        # Build the frontend application
        subprocess.run(['npm', 'run', 'build'])
        
        # Return to the project root
        os.chdir('../')
    except Exception as e:
        print(f"Error building frontend: {e}")

def run_frontend():
    """Runs the frontend application."""
    try:
        # Navigate to the frontend directory
        os.chdir('./frontend')
        
        # Start the development server
        subprocess.run(['npm', 'run', 'start'])
        
        # Return to the project root
        os.chdir('../')
    except Exception as e:
        print(f"Error running frontend: {e}")

def main():
    """Main function."""
    print("Welcome to the frontend-app project!")
    
    while True:
        print("\nOptions:")
        print("1. Build frontend")
        print("2. Run frontend")
        print("3. Quit")
        
        choice = input("Choose an option: ")
        
        if choice == "1":
            build_frontend()
        elif choice == "2":
            run_frontend()
        elif choice == "3":
            break
        else:
            print("Invalid option. Please choose again.")

if __name__ == "__main__":
    main()
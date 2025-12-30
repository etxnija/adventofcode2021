import timeit
import sys
import os
from pathlib import Path

# --- Path Configuration ---
current_file = Path(__file__).resolve()
day_folder = current_file.parent
project_root = day_folder.parent
input_file = day_folder / "input.txt"

# Add project root and day folder to sys.path
sys.path.insert(0, str(project_root))
sys.path.insert(0, str(day_folder))

try:
    from read_data import read_input
except ImportError:
    print(f"Error: Could not import read_input from {project_root}/read_data.py")
    sys.exit(1)

def run_benchmark():
    # Load data using your specific path
    data = read_input(str(input_file))
    
    if not data:
        print(f"Error: No data found in {input_file}")
        return

    try:
        import mojo.importer
        import kernel
        # Note: both versions use find_invalid_py as the entry point
        func = kernel.joltage_py
    except (ImportError, AttributeError) as e:
        print(f"Mojo Error: {e}")
        return

    print(f"--- Benchmarking Day 2 Logic ({len(data)} lines) ---")
    
    def test_logic():
        func(data)

    iterations = 3
    print(f"Running {iterations} iterations on input: {input_file.name}")
    
    total_time = timeit.timeit(test_logic, number=iterations)
    avg_time = total_time / iterations

    print("-" * 40)
    print(f"Average Execution Time: {avg_time:.6f} seconds")
    print("-" * 40)

if __name__ == "__main__":
    run_benchmark()

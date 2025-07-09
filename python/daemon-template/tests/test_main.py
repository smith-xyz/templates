"""
Unit tests for the main module
"""

import pytest
import sys
import os
from io import StringIO

# Add the project root to the Python path
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..'))

from main import main


def test_main_function(capsys):
    """Test that the main function prints expected output."""
    main()
    captured = capsys.readouterr()
    
    assert "Hello, World!" in captured.out
    assert "Welcome to your basic Python project!" in captured.out


def test_main_function_output_format(capsys):
    """Test that the main function output is properly formatted."""
    main()
    captured = capsys.readouterr()
    
    lines = captured.out.strip().split('\n')
    assert len(lines) == 2
    assert lines[0] == "Hello, World!"
    assert lines[1] == "Welcome to your basic Python project!"


class TestMainModule:
    """Test class for main module functionality."""
    
    def test_main_is_callable(self):
        """Test that main function is callable."""
        assert callable(main)
        
    def test_main_runs_without_error(self):
        """Test that main function runs without raising exceptions."""
        try:
            main()
        except Exception as e:
            pytest.fail(f"main() raised an exception: {e}")


if __name__ == "__main__":
    pytest.main([__file__]) 
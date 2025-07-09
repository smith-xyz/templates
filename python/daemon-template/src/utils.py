"""
Utility functions for the basic Python project
"""

import os
import sys
from typing import Optional


def get_project_root() -> str:
    """
    Get the root directory of the project.
    
    Returns:
        str: The absolute path to the project root directory
    """
    current_dir = os.path.dirname(os.path.abspath(__file__))
    return os.path.dirname(current_dir)


def validate_environment() -> bool:
    """
    Validate that the environment is set up correctly.
    
    Returns:
        bool: True if environment is valid, False otherwise
    """
    python_version = sys.version_info
    if python_version.major < 3 or (python_version.major == 3 and python_version.minor < 8):
        print(f"Error: Python 3.8 or higher is required. Current version: {python_version.major}.{python_version.minor}")
        return False
    
    return True


def format_message(message: str, prefix: Optional[str] = None) -> str:
    """
    Format a message with an optional prefix.
    
    Args:
        message (str): The message to format
        prefix (str, optional): Optional prefix to add to the message
        
    Returns:
        str: The formatted message
    """
    if prefix:
        return f"[{prefix}] {message}"
    return message


def safe_get_env(key: str, default: Optional[str] = None) -> Optional[str]:
    """
    Safely get an environment variable.
    
    Args:
        key (str): The environment variable key
        default (str, optional): Default value if key is not found
        
    Returns:
        str or None: The environment variable value or default
    """
    return os.environ.get(key, default) 